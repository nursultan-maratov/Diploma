package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/nursultan-maratov/Diploma.git/internal/repository"
	"strconv"
)

type Middleware struct {
	user repository.UserSDK
}

func NewMiddleware(user repository.UserSDK) *Middleware {
	return &Middleware{
		user: user,
	}
}
func (m *Middleware) SetUserToContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ID := c.Get("id")
		if ID == nil {
			c.Set("id", "empty")
			return next(c)
		}

		StrID := fmt.Sprintf("%v", ID)
		uintID, err := strconv.ParseUint(StrID, 10, 64)
		if err != nil {
			c.Set("id", "not uint")
			return next(c)
		}

		user, err := m.user.GetUser(uint(uintID))
		if err != nil {
			c.Set("id", "empty")
			return next(c)
		}
		c.Set("id", user.ID)

		return next(c)
	}
}
