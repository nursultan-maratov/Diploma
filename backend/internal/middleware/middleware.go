package middleware

import (
	"github.com/nursultan-maratov/Diploma.git/internal/repository"
)

type Middleware struct {
	user repository.UserSDK
}

func NewMiddleware(user repository.UserSDK) *Middleware {
	return &Middleware{
		user: user,
	}
}
