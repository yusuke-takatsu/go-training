package vo

import (
	"errors"
	"regexp"
)

type Email struct {
	address string
}

// NewEmail Because it is valueObject, it is returned as a value type.
func NewEmail(address string) (Email, error) {
	if !isValidEmail(address) {
		return Email{}, errors.New("invalid Email Address")
	}

	return Email{address: address}, nil
}

func isValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return regex.MatchString(email)
}

func (e Email) Value() string {
	return e.address
}
