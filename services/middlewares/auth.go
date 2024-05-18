package middlewares

import (
	"net/http"
	"project-auction/services/servicesmodels"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"project-auction/apperrors"
	"project-auction/services"
)

func ParseAccessToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error

		var customCtx *servicesmodels.CustomContext
		if ctx, ok := c.(*servicesmodels.CustomContext); ok {
			customCtx = ctx
		} else {
			customCtx = &servicesmodels.CustomContext{Context: c}
		}

		accessToken := c.Request().Header.Get("Authorization")

		customCtx.SellerID, err = services.ParseJWTAccessToken(accessToken)
		if err != nil {
			log.Error("failed to parse access token", err)
			return c.JSON(http.StatusUnauthorized, apperrors.NewAuthorization("failed to parse access token"))
		}

		return next(customCtx)
	}
}
