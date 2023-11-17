package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Login struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (login *Login) TrimEmail(email string) error {
	login.Email = html.EscapeString(strings.TrimSpace(strings.Join(strings.Fields(email), "")))

	return nil
}

func (login *Login) HashPwd(pwd string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 16)

	if err != nil {
		return err
	}

	login.Password = string(bytes)

	return nil
}

func (login *Login) CheckPwd(pwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(pwd))

	if err != nil {
		return err
	}

	return nil
}
