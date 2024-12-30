package vo

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type Password struct {
	password string
}

// NewPassword Because it is valueObject, it is returned as a value type.
func NewPassword(password string) (Password, error) {
	if len(password) < 8 || len(password) > 20 {
		return Password{}, errors.New("password must be between 8 and 20 characters long")
	}

	if !isValidPassword(password) {
		return Password{}, errors.New("password must contain at least one uppercase letter, one number, and only alphanumeric characters")
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return Password{}, err
	}
	return Password{password: hashedPassword}, nil
}

func isValidPassword(password string) bool {
	alphanumericRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if matched := alphanumericRegex.MatchString(password); !matched {
		return false
	}

	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	if matched := uppercaseRegex.MatchString(password); !matched {
		return false
	}

	numberRegex := regexp.MustCompile(`[0-9]`)
	if matched := numberRegex.MatchString(password); !matched {
		return false
	}

	return true
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (p Password) Value() string {
	return p.password
}
