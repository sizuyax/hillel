package dto

import "github.com/labstack/echo/v4"

type CustomContext struct {
	echo.Context
	ProfileID   int
	ProfileType string
}
