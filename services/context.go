package services

import (
	"context"
	"github.com/labstack/echo/v4"
	"project-auction/models"
	"project-auction/services/servicesmodels"
)

func NewContextFromEchoContext(c echo.Context) (context.Context, error) {
	var customCtx *servicesmodels.CustomContext
	if ctx, ok := c.(*servicesmodels.CustomContext); ok {
		customCtx = ctx
	} else {
		cc := &servicesmodels.CustomContext{Context: c}
		customCtx = cc
	}

	ctx := context.WithValue(context.Background(), models.SellerIDKey, customCtx.SellerID)

	return ctx, nil
}
