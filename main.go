package main

import (
	"github.com/mddsilva/gobet/internal/router"
	"github.com/mddsilva/gobet/pkg/config"
)

func main() {

	init := config.Init()
	server := router.Init(init)

	server.Run(":8080")
}
