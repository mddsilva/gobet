// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

import (
	"github.com/google/wire"
	"github.com/mddsilva/gobet/internal/auth/controller"
	"github.com/mddsilva/gobet/internal/auth/repository"
	"github.com/mddsilva/gobet/internal/auth/service"
)

// Injectors from wire.go:

func Init() *Initialization {
	gormDB := Connect()
	userRepositoryImpl := repository.UserRepositoryInit(gormDB)
	authServiceImpl := service.AuthServiceInit(userRepositoryImpl)
	authControllerImpl := controller.AuthControllerInit(authServiceImpl)
	initialization := NewInitialization(userRepositoryImpl, authServiceImpl, authControllerImpl)
	return initialization
}

// wire.go:

var db = wire.NewSet(Connect)

var authServiceSet = wire.NewSet(service.AuthServiceInit, wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)))

var userRepoSet = wire.NewSet(repository.UserRepositoryInit, wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)))

var authCtrlSet = wire.NewSet(controller.AuthControllerInit, wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)))
