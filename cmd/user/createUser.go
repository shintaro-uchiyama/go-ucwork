package main

import (
	"fmt"
	application "github.com/shintaro-uchiyama/go-ucwork/pkg/application/user"
	infrastructure "github.com/shintaro-uchiyama/go-ucwork/pkg/infrastructure/inMemory/user"
	presentation "github.com/shintaro-uchiyama/go-ucwork/pkg/presentation/cmd"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.WarnLevel)
}

func main()  {
	controller := presentation.NewUserController(
		application.NewUserApplicationService(
			infrastructure.NewInMemoryUserRepository(),
		),
	)
	user, err := controller.Regist()
	if err != nil {
		log.Error("regist fail: ", err)
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("regist success: %+v", user))
}
