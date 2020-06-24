package email

import (
	"errors"
)

type Email struct{
	value string
}

func New(email string) (*Email, error) {
	
	if len(email) <= 0 {
		return nil, errors.New("Invalid email address")
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

