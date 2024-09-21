// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"github.com/mddsilva/gobet/internal/auth/controller"
	"github.com/mddsilva/gobet/internal/auth/repository"
	"github.com/mddsilva/gobet/internal/auth/service"
)

var db = wire.NewSet(Connect)

var authServiceSet = wire.NewSet(service.AuthServiceInit, wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)))

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

var authCtrlSet = wire.NewSet(controller.AuthControllerInit, wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)))

func Init() *Initialization {
	wire.Build(NewInitialization, db, authCtrlSet, authServiceSet, userRepoSet)
	return nil
}
