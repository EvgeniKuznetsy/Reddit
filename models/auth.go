package models

import (
	"errors"
	"fmt"
	"strings"
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

	if strings.Contains(i.Login, "@") {
		return errors.New(`login have forbidden symbol "@"`)
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
	Identifier string
	Password   string
}

func (i *InputSingIn) Validate() error {
	if i.Identifier == "" {
		return errors.New("login required")
	}

	if i.Password == "" {
		return errors.New("password required")
	}

	return nil
}

type OutPutIn struct {
	Session string `json:"session"`
	Account string `json:"	Account"`
}
