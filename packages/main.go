package main

import (
	"fmt"

	"github.com/Aytaditya/go-learn/auth"
	"github.com/Aytaditya/go-learn/user"
)

func main() {
	bool := auth.LoginWithCredentials("admin", "password") // function name should be in capital letters to be exported
	if bool {
		println("Login Successful")
	} else {
		println("Login Failed")
	}

	fmt.Println(auth.SessionInfo())
	vale := user.GetUserDetails()
	fmt.Println(vale)
}
