package services

import (
	"errors"
	jwtauth "gokitdemo/auth"
)

func (s BasicService) Login(name, pwd string) (string, error) {
	if name == "name" && pwd == "pwd" {
		token, err := jwtauth.Sign(name, pwd)
		return token, err
	}

	return "", errors.New("Your name or password dismatch")
}
