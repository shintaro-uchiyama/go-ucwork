package main

import (
	"fmt"
	"os"

	"github.com/shintaro-uchiyama/go-ucwork/pkg/application"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/infrastructure"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/presentation"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	f, err := os.OpenFile("/var/log/ucwork.log", os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(fmt.Errorf("[main.init] %w", err))
	}
	log.SetOutput(f)
	log.SetLevel(log.InfoLevel)
}

func main() {
	controller := presentation.NewUserController(
		application.NewUserApplicationService(
			infrastructure.NewInMemoryUserRepository(),
			domain.NewUserService(infrastructure.NewInMemoryUserRepository()),
		),
	)
	user, err := controller.Create()
	if err != nil {
		log.Error("Create fail: ", err)
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("Create success: %+v", user))
}
