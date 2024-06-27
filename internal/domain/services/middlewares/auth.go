package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"project-auction/internal/common/apperrors"
	"project-auction/internal/domain/services"
	"project-auction/internal/domain/services/dto"
)

const (
	authorization = "Authorization"
)

func ParseAccessUserToken(tokenService services.TokenService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var err error

			var customCtx *dto.CustomContext
			if ctx, ok := c.(*dto.CustomContext); ok {
				customCtx = ctx
			} else {
				customCtx = &dto.CustomContext{Context: c}
			}

			accessToken := c.Request().Header.Get(authorization)

			customCtx.ProfileID, customCtx.ProfileType, err = tokenService.ParseJWTAccessToken(accessToken)
			if err != nil {
				log.Error("failed to parse access token", err)
				return c.JSON(http.StatusUnauthorized, apperrors.NewAuthorization("failed to parse access token"))
			}

			return next(customCtx)
		}
	}
}

func ParseAccessSellerToken(tokenService services.TokenService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var err error

			var customCtx *dto.CustomContext
			if ctx, ok := c.(*dto.CustomContext); ok {
				customCtx = ctx
			} else {
				customCtx = &dto.CustomContext{Context: c}
			}

			accessToken := c.Request().Header.Get(authorization)

			customCtx.ProfileID, customCtx.ProfileType, err = tokenService.ParseJWTAccessToken(accessToken)
			if err != nil {
				log.Error("failed to parse access token", err)
				return c.JSON(http.StatusUnauthorized, apperrors.NewAuthorization("failed to parse access token"))
			}

			if customCtx.ProfileType != "seller" {
				return c.JSON(http.StatusForbidden, apperrors.NewAuthorization("user isn't seller"))
			}

			return next(customCtx)
		}
	}
}
