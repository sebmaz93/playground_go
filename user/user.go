package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User     User // you use it anonymously also "User"
}

func (u *User) OutputUserData() {
	// here u.firstName is the shortcut to (*u).firstName
	fmt.Println(u.firstName, u.lastName)
}

func (u *User) ClearUserData() {
	u.firstName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email, password, User{
			"Admin",
			"Admin",
			"---",
			time.Now(),
		},
	}
}

func New(firstName, lastName, birthday string) (*User, error) {
	if firstName == "" || lastName == "" {
		return nil, errors.New("firstname lastname are required")
	}

	return &User{
		firstName,
		lastName,
		birthday,
		time.Now(),
	}, nil
}
