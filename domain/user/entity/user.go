package entity

import (
	"github.com/yusuke-takatsu/go-training/domain/user/vo"
	"github.com/yusuke-takatsu/go-training/enum"
	"github.com/yusuke-takatsu/go-training/util"
	"time"
)

type User struct {
	ID        string
	Email     vo.Email
	Password  vo.Password
	Image     string
	Status    enum.Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

func newUser(
	id string,
	email vo.Email,
	password vo.Password,
	image string,
	status enum.Status,
) *User {
	now := time.Now()

	return &User{
		ID:        id,
		Email:     email,
		Password:  password,
		Image:     image,
		Status:    status,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func CreateUser(
	email vo.Email,
	password vo.Password,
	image string,
) *User {
	return newUser(
		util.GenerateIdentifier().Identifier,
		email,
		password,
		image,
		enum.TemporaryMember,
	)
}
