package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/nursultan-maratov/Diploma.git/internal/repository"
	"github.com/uptrace/bun"
	"net/http"
	"strings"
	"time"
)

type Middleware struct {
	user   repository.UserSDK
	secret string
}

func NewMiddleware(db bun.IDB, secret string) *Middleware {
	return &Middleware{
		user:   repository.NewUserRepo(db),
		secret: secret,
	}
}

func (m *Middleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing or invalid token")
		}

		claims, err := m.ValidateToken(tokenString)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		userIDFloat, ok := claims["user_id"].(float64) // JWT хранит числа как float64
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid user_id in token")
		}
		userID := uint(userIDFloat) // Приводим к uint

		userInfo, err := m.user.GetUser(c.Request().Context(), userID)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "user not found")
		}

		c.Set("user_id", userID)
		c.Set("user", userInfo)

		return next(c)
	}
}

func (m *Middleware) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(m.secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

func (m *Middleware) GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secret))
}
