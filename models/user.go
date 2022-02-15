package models

import "net/mail"

type User struct {
	Id    string
	Name  string `json:"name" binding:"required"`
	Email string `json:"Email" binding:"required"`
}

func (u *User) Validate() (bool, string) {
	if len(u.Name) < 3 {
		return false, "username less then 3 characters"
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return false, "not valid email address"
	}

	return true, ""
}
