package api

import (
	"errors"
	"strings"

	"github.com/JueViGrace/closs-server-local/internal/data"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/JueViGrace/closs-server-local/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (a *api) sessionMiddleware(c *fiber.Ctx) error {
	_, err := getUserDataForReq(c, a.db)
	if err != nil {
		res := types.RespondUnauthorized(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

func (a *api) authenticatedHandler(handler types.AuthDataHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data, err := getUserDataForReq(c, a.db)
		if err != nil {
			res := types.RespondUnauthorized(nil, err.Error())
			return c.Status(res.Status).JSON(res)
		}

		return handler(c, data)
	}
}

func getUserDataForReq(c *fiber.Ctx, db data.Storage) (*types.AuthData, error) {
	jwt, err := extractJWTFromHeader(c, func(s string) {
		db.SessionStorage().DeleteSessionByToken(s)
	})
	if err != nil {
		return nil, err
	}

	session, err := db.SessionStorage().GetSessionById(jwt.Claims.UserId)
	if err != nil {
		return nil, err
	}

	dbUser, err := db.UserStore().GetDbUserByUsername(session.Username)
	if err != nil {
		return nil, err
	}

	// TODO: fill code
	return &types.AuthData{
		UserId:   session.UserId,
		Username: dbUser.Username,
		Code:     "",
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
		log.Info(err.Error())
		expired(tokenString)
		return nil, errors.New("permission denied")
	}

	claims, ok := token.Claims.(*util.JWTClaims)
	if !ok || !token.Valid {
		log.Info("invalid claims or expired")
		expired(tokenString)
		return nil, errors.New("permission denied")
	}

	if len(claims.Audience) > 1 || claims.
		Audience[0] != "api" {
		log.Infof("bad audience: %v", claims.Audience)
		expired(tokenString)
		return nil, errors.New("permision denied")
	}

	if claims.Issuer != util.Issuer {
		log.Infof("bad issuer: %v", claims.Issuer)
		expired(tokenString)
		return nil, errors.New("permision denied")
	}

	return &types.JwtData{
		Token:  token,
		Claims: claims,
	}, nil
}
