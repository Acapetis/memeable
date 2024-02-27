package handler

import (
	"github.com/Acapetis/whatDoUMeme/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service service.IService
}

func NewHandler(s service.IService) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) GetAll(ctx echo.Context) error {
	result := h.service.GetAll(ctx)
	return ctx.JSON(200, result)
}
