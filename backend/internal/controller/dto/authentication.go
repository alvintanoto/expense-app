package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type LoginRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (dto LoginRequestDTO) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Username, validation.Required, validation.Length(5, 30), is.Alphanumeric),
		validation.Field(&dto.Password, validation.Required, validation.Length(8, 30), is.Alphanumeric),
	)
}

type LoginResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenDTO struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
