package handlers

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"project-auction/internal/common/apperrors"
	"project-auction/internal/domain/entity"
	"project-auction/internal/domain/services"
	"strconv"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h Handler) WebSocket(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
	if err != nil {
		h.log.Error("failed to upgrade websocket", slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}
	defer conn.Close()

	ctx, err := services.NewContextFromEchoContext(c)
	if err != nil {
		h.log.Error("failed get context from echo context", slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	itemIDString := c.Param("id")

	ownerID, ok := ctx.Value(entity.ProfileIDKey).(int)
	if !ok {
		h.log.Error("failed to convert owner id to int")
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	ownerIDString := strconv.Itoa(ownerID)

	fileName := ownerIDString + "_bids_" + itemIDString + ".json"

	obj, err := h.minioStorage.Object(fileName)
	if err != nil {
		h.log.Error("failed to ")
		return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
	}

	for {
		if err := conn.WriteMessage(websocket.TextMessage, obj); err != nil {
			h.log.Error("failed to write message", slog.String("error", err.Error()))
			return c.JSON(http.StatusInternalServerError, apperrors.NewInternal())
		}

		time.Sleep(1 * time.Minute)
	}
}
