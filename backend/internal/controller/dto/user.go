package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UserWalletRequestDTO struct {
	WalletName     string `json:"wallet_name"`
	CurrencyID     string `json:"currency_id"`
	InitialBalance string `json:"initial_balance"`
}

func (dto UserWalletRequestDTO) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.WalletName, validation.Required, validation.Length(3, 20), is.ASCII),
		validation.Field(&dto.CurrencyID, validation.Required),
		validation.Field(&dto.InitialBalance, validation.Required, validation.Length(0, 16), is.Digit),
	)
}

type UserWalletDTO struct {
	WalletID   string `json:"wallet_id"`
	WalletName string `json:"wallet_name"`
}

type UserWalletCurrencyDTO struct {
	CurrencyID   string `json:"currency_id"`
	CurrencyName string `json:"currency_name"`
	CurrencyCode string `json:"currency_code"`
}

type UserWalletResponseDTO struct {
	UserWallet UserWalletDTO         `json:"wallet"`
	Currency   UserWalletCurrencyDTO `json:"currency"`
}
