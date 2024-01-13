package main

import (
	"fmt"

	"example.com/user/userOps"
)

func getUserData(infoText string) string {
	var userInput string
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birth date: ")
	fmt.Print(firstName)
	fmt.Print(lastName)

	var appUser *userOps.User

	appUser, err := userOps.New(firstName, lastName, birthDate)

	if err != nil {
		fmt.Printf("Error occured while creating user")
		fmt.Print(err)
		return
	}

	admin := userOps.NewAdmin("stanley", "ds")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}
