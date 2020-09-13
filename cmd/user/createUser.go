package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.WarnLevel)
}

func main()  {
	fmt.Print("Please Enter email: ")
	var email string
	_, emailErr := fmt.Scanln(&email)
	if emailErr != nil {
		log.Error(emailErr)
		os.Exit(1)
	}
	fmt.Println(email)
	fmt.Println(emailErr)

	fmt.Print("Please Enter password: ")
	var password string
	_, passwordErr := fmt.Scanln(&password)
	if passwordErr != nil {
		log.Error(passwordErr)
		os.Exit(1)
	}
	fmt.Println(password)
	fmt.Println(passwordErr)
}
