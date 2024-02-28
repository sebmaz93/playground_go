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

//type Text struct {
//	Content string `json:"firstname"` // this will make json package read by this key its called struct tag
//}
//
//func Save() {
//	fileName := strings.ReplaceAll("some string", " ", "_") // this will replace all white space with underscore
//	fileName := strings.ToLower("QWERTY")
//
//	json := json.Marshal("seb maz") // this will only write the data that is public (Capital Case in struct)
//	os.WriteFile('fileName', json, 0644)
//
//	bufio.NewReader()
//}
