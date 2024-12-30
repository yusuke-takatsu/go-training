package usecase

import (
	"github.com/yusuke-takatsu/go-training/domain/user/entity"
	"github.com/yusuke-takatsu/go-training/domain/user/vo"
	"github.com/yusuke-takatsu/go-training/exception"
	"github.com/yusuke-takatsu/go-training/infra/user/repository"
	"github.com/yusuke-takatsu/go-training/interface/user/dto"
)

type UseCase struct {
	repo repository.Repository
}

func NewUserUseCase(repo repository.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (s *UseCase) Register(input dto.RegisterInput) error {
	emailVo, err := vo.NewEmail(input.Email)
	if err != nil {
		return exception.InValid.Wrap(err, err.Error())
	}

	passwordVo, err := vo.NewPassword(input.Password)
	if err != nil {
		return exception.InValid.Wrap(err, err.Error())
	}

	userEntity := entity.CreateUser(emailVo, passwordVo, input.Image)
	if err := s.repo.Save(userEntity); err != nil {
		return exception.InsertFailed.Wrap(err, "fail to record data")
	}

	return nil
}
