package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}
func (h *Handler) HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
