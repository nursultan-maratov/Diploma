package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/nursultan-maratov/Diploma.git/internal/manager/order"
	"github.com/nursultan-maratov/Diploma.git/internal/manager/user"
	"github.com/nursultan-maratov/Diploma.git/internal/model"
	"log"
	"time"

	"net/http"
)

type Handler struct {
	userManager    user.ManagerSDK
	productManager order.ManagerSDK
	secret         string
}

func NewHandler(userManager user.ManagerSDK, productManager order.ManagerSDK, secret string) *Handler {
	return &Handler{
		userManager:    userManager,
		productManager: productManager,
		secret:         secret,
	}
}

func (h *Handler) CreateUsers(c echo.Context) error {
	var req *model.UserRequest
	err := c.Bind(&req)
	log.Println(err)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	id, err := h.userManager.CreateUser(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, id)
}

func (h *Handler) Auth(c echo.Context) error {
	var req model.Auth
	err := c.Bind(&req)

	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	isHash, id, err := h.userManager.Auth(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}

	if isHash != true {

	}
	token, err := h.generateToken(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not generate token")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func (h *Handler) BuyProduct(c echo.Context) error {
	var req *model.BuyProduct
	err := c.Bind(&req)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	req.UserID = userID

	resp, err := h.productManager.BuyProduct(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) ListOrder(c echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	resp, err := h.productManager.ListOrder(c.Request().Context(), int(userID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	if resp == nil {
		return c.JSON(http.StatusOK, "Список пуст")

	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) ListProduct(c echo.Context) error {
	resp, err := h.productManager.ListProduct(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	if resp == nil {
		return c.JSON(http.StatusOK, "Список пуст")

	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) generateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(h.secret))
}
