package api

import (
	"errors"
	"strings"
	"time"

	"github.com/JueViGrace/closs-server-local/internal/data"
	"github.com/JueViGrace/closs-server-local/internal/types"
	"github.com/JueViGrace/closs-server-local/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
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

func getUserDataForReq(c *fiber.Ctx, db *data.Storage) (*types.AuthData, error) {
	jwt, err := extractJWTFromHeader(c, func(s uuid.UUID) {
		db.CacheStore.SessionStorage().DeleteSession(s)
	})
	if err != nil {
		return nil, err
	}

	session, err := db.CacheStore.SessionStorage().GetSessionById(jwt.Claims.UserId)
	if err != nil {
		return nil, err
	}

	dbUser, err := db.MyStore.Queries.GetUserByUsername(db.MyStore.Ctx, session.Username)
	if err != nil {
		return nil, err
	}

	// todo: fill code
	return &types.AuthData{
		UserId:   session.UserId,
		Username: dbUser.Username,
		Code:     "",
	}, nil
}

func extractJWTFromHeader(c *fiber.Ctx, expired func(uuid.UUID)) (*types.JwtData, error) {
	header := strings.Join(c.GetReqHeaders()["Authorization"], "")

	if !strings.HasPrefix(header, "Bearer") {
		return nil, errors.New("permission denied")
	}

	tokenString := strings.Split(header, " ")[1]
	token, err := util.ValidateJWT(tokenString)
	if err != nil {
		log.Info(err.Error())
		return nil, errors.New("permission denied")
	}

	if !token.Valid {
		log.Info("invalid token")
		return nil, errors.New("permission denied")
	}

	claims, ok := token.Claims.(*util.JWTClaims)
	if !ok {
		log.Info("invalid claims")
		return nil, errors.New("permission denied")
	}

	if claims.ExpiresAt.Time.Unix() < time.Now().Unix() {
		log.Info("expired")
		expired(claims.UserId)
		return nil, errors.New("permision denied")
	}

	if len(claims.Audience) > 1 || claims.
		Audience[0] != "api" {
		log.Infof("bad audience: %v", claims.Audience)
		return nil, errors.New("permision denied")
	}

	if claims.Issuer != util.Issuer {
		log.Infof("bad issuer: %v", claims.Issuer)
		return nil, errors.New("permision denied")
	}

	return &types.JwtData{
		Token:  token,
		Claims: *claims,
	}, nil
}
