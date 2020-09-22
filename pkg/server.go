package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/application"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/infrastructure"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/presentation"
	"github.com/sirupsen/logrus"
)

func initRoute() {
	userHandler := presentation.NewUserHandler(
		application.NewUserApplicationService(
			infrastructure.NewInMemoryUserRepository(),
			domain.NewUserService(
				infrastructure.NewInMemoryUserRepository(),
			),
		),
	)
	r := gin.Default()
	r.Use(logMiddleWare())
	r.POST("/users", userHandler.Create)
	r.GET("/users", userHandler.List)
	r.Run()
}

func logMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := logrus.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"url": c.Request.URL,
		})
		logger.Info("start")
		c.Next()
		logger.Info("end")
	}
}

func main() {
	initRoute()
}
