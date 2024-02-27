package main

import (
	"fmt"
	"playground/user"
)

func main() {
	userFirstName := getUserData("please enter your name: ")
	userLastName := getUserData("please enter your lastname: ")
	userBirthday := getUserData("please enter your birthday (DD/MM/YYYY): ")

	var user1 *user.User

	// if you don't define the keys then the order must be the same as struct
	user1, err := user.New(userFirstName, userLastName, userBirthday)

	if err != nil {
		return
	}

	admin := user.NewAdmin("seb@mail.com", "1234")

	// if we use anonymous User on admin struct we can access the methods directly from the Admin
	admin.User.OutputUserData()

	user1.OutputUserData()
	user1.ClearUserData()
	user1.OutputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	_, err := fmt.Scan(&value)
	if err != nil {
		return ""
	}
	return value
}
