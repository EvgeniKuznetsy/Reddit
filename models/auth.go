package models

import (
	"errors"
	"fmt"
)

type InputSinUp struct {
	Login    string `binding:"required"`
	Password string `binding:"required"`
	Name     string `binding:"required"`
	Email    string `binding:"required"`
}

const MinPasswordLength = 8
const MinLoginLength = 6

func (i *InputSinUp) IsValid() error {
	if len(i.Login) < MinLoginLength {
		return errors.New(fmt.Sprintf("login min length %d symbols", MinPasswordLength))
	}

	if len(i.Password) < MinPasswordLength {
		return errors.New(fmt.Sprintf("password min length %d symbols", MinPasswordLength))
	}

	return nil
}

type OutPutUp struct {
	Session string `json:"session"`
}
type InputSingIn struct {
	Login    string
	Password string
}
type OutPutIn struct {
	Session string `json:"session"`
	Account string `json:"account"`
}
