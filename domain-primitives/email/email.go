package email

import (
	"errors"
	"strings"
)


const (
	minLength = 5	
)

var (
	ErrInvalidEmail  = errors.New("Invalid Email.")
	ErrEmailTooShort = errors.New("Invalid email: min length allowed is 5")
)


type Email struct{
	value string
}

func New(e string) (*Email, error) {
	
	email := strings.TrimSpace(e) 
	if len(email) <= 0 {
		return nil, ErrInvalidEmail
	}

	if len(email) <= minLength {
		return nil, ErrEmailTooShort
	}

	return &Email{
		value: email,
	}, nil
}

func (e Email) Value() string {
	return e.value
}

func (e Email) String() string {
	return e.value
}

func (e Email) Equals(other Email) bool {
	if e.value == other.value {
		return true
	}
	
	return false
}

