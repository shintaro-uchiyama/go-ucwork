package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Print("Please Enter email: ")
	var email string
	_, emailErr := fmt.Scanln(&email)
	if emailErr != nil {
		fmt.Fprintln(os.Stderr, emailErr)
		os.Exit(1)
	}
	fmt.Println(email)
	fmt.Println(emailErr)

	fmt.Print("Please Enter password: ")
	var password string
	_, passwordErr := fmt.Scanln(&password)
	if passwordErr != nil {
		fmt.Fprintln(os.Stderr, emailErr)
		os.Exit(1)
	}
	fmt.Println(password)
	fmt.Println(passwordErr)
}
