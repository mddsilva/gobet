package config

import (
	"github.com/mddsilva/gobet/internal/auth/controller"
	"github.com/mddsilva/gobet/internal/auth/repository"
	"github.com/mddsilva/gobet/internal/auth/service"
)

type Initialization struct {
	userRepo repository.UserRepository
	authSvc  service.AuthService
	AuthCtrl controller.AuthController
}

func NewInitialization(userRepo repository.UserRepository,
	userService service.AuthService,
	authCtrl controller.AuthController) *Initialization {
	return &Initialization{
		userRepo: userRepo,
		authSvc:  userService,
		AuthCtrl: authCtrl,
	}
}
