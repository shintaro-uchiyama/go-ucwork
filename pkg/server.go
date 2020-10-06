package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/application"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/infrastructure"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/presentation"
	"github.com/sirupsen/logrus"
)

func initRoute() {
	userFirebaseAuth, firebaseErr := infrastructure.NewUserFirebaseAuth()
	if firebaseErr != nil {
		logrus.Fatal(fmt.Errorf("[main.initRoute]: %w", firebaseErr))
		panic(firebaseErr)
	}
	inMemoryUserRepository := infrastructure.NewInMemoryUserRepository()
	userHandler := presentation.NewUserHandler(
		application.NewUserApplicationService(
			inMemoryUserRepository,
			userFirebaseAuth,
		),
	)
	r := gin.Default()
	r.Use(logMiddleWare())
	r.POST("/users", userHandler.Create)
	r.GET("/users", userHandler.List)
	err := r.Run()
	if err != nil {
		logrus.Fatal(fmt.Errorf("[main.initRoute]: %w", err))
	}
}

func logMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := logrus.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"url":    c.Request.URL,
		})
		logger.Info("start")
		c.Next()
		logger.Info("end")
	}
}

func main() {
	initRoute()
}
