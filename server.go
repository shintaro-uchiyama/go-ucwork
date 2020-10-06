package main

import (
	"fmt"

	"github.com/shintaro-uchiyama/go-ucwork/pkg/application"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/infrastructure"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/shintaro-uchiyama/go-ucwork/graph"
	"github.com/shintaro-uchiyama/go-ucwork/graph/generated"
)

func graphqlHandler() gin.HandlerFunc {
	inMemoryUserRepository := infrastructure.NewInMemoryUserRepository()
	userFirebaseAuth, firebaseErr := infrastructure.NewUserFirebaseAuth()
	if firebaseErr != nil {
		logrus.Fatal(fmt.Errorf("[main.initRoute]: %w", firebaseErr))
		panic(firebaseErr)
	}
	resolver := graph.NewResolver(
		application.NewUserApplicationService(
			inMemoryUserRepository,
			userFirebaseAuth,
		),
	)
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	r := gin.Default()
	r.Use(logMiddleWare())
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
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
