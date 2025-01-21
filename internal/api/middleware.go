package api

import (
	"errors"
	"strings"
	"time"

	"github.com/JueViGrace/closs-server-local/internal/data"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/JueViGrace/closs-server-local/internal/util"
	"github.com/gofiber/fiber/v2"
)

func (a *api) sessionMiddleware(c *fiber.Ctx) error {
	data, err := getUserDataForReq(c, a.db)
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	_, err = a.db.Queries.GetSessionById(a.db.Ctx, data.Jwt.Claims.UserId.String())
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

func (a *api) authenticatedRoute(handler types.AuthDataHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data, err := getUserDataForReq(c, a.db)
		if err != nil {
			res := types.RespondBadRequest(err.Error(), "Failure")
			return c.Status(res.Status).JSON(res)
		}

		_, err = a.db.Queries.GetSessionById(a.db.Ctx, data.Jwt.Claims.UserId.String())
		if err != nil {
			res := types.RespondUnauthorized(err.Error(), "Failed")
			return c.Status(res.Status).JSON(res)
		}

		return handler(c, data)
	}
}

func getUserDataForReq(c *fiber.Ctx, db *data.Storage) (*types.AuthData, error) {
	jwt, err := extractJWTFromHeader(c, func(s string) {
		db.Queries.DeleteSessionByToken(db.Ctx, s)
	})
	if err != nil {
		return nil, err
	}

	// user, err := db.Queries.GetUserById(db.Ctx, jwt.Claims.UserId.String())
	// if err != nil {
	// 	return nil, err
	// }

	return &types.AuthData{
		Jwt: *jwt,
	}, nil
}

func extractJWTFromHeader(c *fiber.Ctx, expired func(string)) (*types.JwtData, error) {
	header := strings.Join(c.GetReqHeaders()["Authorization"], "")

	if !strings.HasPrefix(header, "Bearer") {
		return nil, errors.New("permission denied")
	}

	tokenString := strings.Split(header, " ")[1]
	token, err := util.ValidateJWT(tokenString)
	if err != nil {
		return nil, errors.New("permission denied")
	}

	if !token.Valid {
		expired(tokenString)
		return nil, errors.New("permission denied")
	}

	claims, ok := token.Claims.(util.JWTClaims)
	if !ok {
		expired(tokenString)
		return nil, errors.New("permission denied")
	}

	if claims.ExpiresAt.Time.UTC().Unix() < time.Now().UTC().Unix() {
		expired(tokenString)
		return nil, errors.New("permision denied")
	}

	if len(claims.Audience) > 1 || claims.
		Audience[0] != "api" {
		return nil, errors.New("permision denied")
	}

	if claims.Issuer != util.Issuer {
		return nil, errors.New("permision denied")
	}

	return &types.JwtData{
		Token:  token,
		Claims: claims,
	}, nil
}
