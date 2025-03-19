package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/nursultan-maratov/Diploma.git/internal/manager/user"
	"github.com/nursultan-maratov/Diploma.git/internal/model"
	"net/http"
)

type Handler struct {
	userManager user.ManagerSDK
}

func NewHandler(userManager user.ManagerSDK) *Handler {
	return &Handler{
		userManager: userManager,
	}
}

func (h *Handler) CreateUsers(c echo.Context) error {
	var req model.UserRequest
	err := c.Bind(&req)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	id, err := h.userManager.CreateUser(c.Request().Context(), &req)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return c.JSON(http.StatusOK, id)
}
