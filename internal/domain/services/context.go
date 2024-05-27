package services

import (
	"context"
	"github.com/labstack/echo/v4"
	"project-auction/internal/domain/entity"
	"project-auction/internal/domain/services/dto"
)

func NewContextFromEchoContext(c echo.Context) (context.Context, error) {
	var customCtx *dto.CustomContext
	if ctx, ok := c.(*dto.CustomContext); ok {
		customCtx = ctx
	} else {
		cc := &dto.CustomContext{Context: c}
		customCtx = cc
	}

	ctx := context.WithValue(context.Background(), entity.ProfileIDKey, customCtx.ProfileID)
	ctx = context.WithValue(ctx, entity.ProfileTypeKey, customCtx.ProfileType)

	return ctx, nil
}
