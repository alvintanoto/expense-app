package service

import (
	"expense_app/internal/entity"
	"expense_app/internal/repository"
	"expense_app/internal/util/logger"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserService interface {
		Register(username, email, password string) error
	}

	implUser struct {
		logger     logger.Logger
		repository repository.Holder
	}
)

func NewUserService(logger logger.Logger, repository repository.Holder) (UserService, error) {
	return &implUser{logger: logger, repository: repository}, nil
}

func (i *implUser) Register(username, email, password string) error {
	var user entity.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		i.logger.Error(err.Error())
		return err
	}

	user.ID = uuid.NewString()
	user.Username = username
	user.Email = email
	user.Password = string(hashedPassword)
	user.CreatedBy = user.ID
	user.IsActive = true

	if err = i.repository.UserRepository.CreateUser(&user); err != nil {
		return err
	}

	return nil
}
