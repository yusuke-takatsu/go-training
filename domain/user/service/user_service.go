package service

import (
	"github.com/yusuke-takatsu/go-training/domain/user/vo"
	"github.com/yusuke-takatsu/go-training/infra/user/repository"
)

type UserService struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) UserService {
	return UserService{repo: repo}
}

func (s UserService) ExistEmail(email vo.Email) (bool, error) {
	exist, err := s.repo.ExistsEmail(email.Value())
	if err != nil {
		return false, err
	}

	return exist, nil
}
