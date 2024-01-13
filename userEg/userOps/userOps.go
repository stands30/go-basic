package userOps

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(fname string, lname string, bdate string) (*User, error) {
	if fname == "" || lname == "" || bdate == "" {
		return nil, errors.New("first name, last name and birthdate cannot be empty")
	}
	return &User{
		firstName: fname,
		lastName:  lname,
		birthDate: bdate,
		createdAt: time.Now(),
	}, nil

}
