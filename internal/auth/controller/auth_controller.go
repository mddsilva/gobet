package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mddsilva/gobet/internal/auth/service"
)

type AuthController interface {
	Signup(c *gin.Context)
}

type AuthControllerImpl struct {
	svc service.AuthService
}

func (u AuthControllerImpl) Signup(c *gin.Context) {
	u.svc.Signup(c)
}

func AuthControllerInit(authService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		svc: authService,
	}
}
