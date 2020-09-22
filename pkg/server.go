package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/application"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/infrastructure"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/presentation"
)

func main() {
	userHandler := presentation.NewUserHandler(
		application.NewUserApplicationService(
			infrastructure.NewInMemoryUserRepository(),
			domain.NewUserService(
				infrastructure.NewInMemoryUserRepository(),
			),
		),
	)
	r := gin.Default()
	r.POST("/users", userHandler.Create)
	r.Run()
}
