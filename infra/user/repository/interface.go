package repository

import (
	"github.com/yusuke-takatsu/go-training/domain/user/entity"
)

type Repository interface {
	Save(user *entity.User) error
}
