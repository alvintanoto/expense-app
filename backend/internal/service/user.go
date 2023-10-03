package service

import (
	"context"
	"expense_app/internal/controller/dto"
	"expense_app/internal/entity"
	"expense_app/internal/repository"
	"expense_app/internal/util/logger"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserService interface {
		Register(username, email, password string) error

		// Create user wallet then after creating wallet this process will populate
		// TODO: categories for the wallet and insert one initial balance data.
		// TODO: limit user wallet to 1 (free tier)/ 3 wallet
		CreateUserWallet(userID, walletName, currencyID, initialBalance string) error
		GetUserWallet(ctx context.Context, userID string) ([]*dto.UserWalletResponseDTO, error)
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

func (i *implUser) CreateUserWallet(userID, walletName, currencyID, initialBalance string) (err error) {
	wallet := &entity.Wallet{
		ID:         uuid.NewString(),
		UserID:     userID,
		CurrencyID: currencyID,
		WalletName: walletName,
		GeneralData: entity.GeneralData{
			CreatedBy: userID,
		},
	}

	if err = i.repository.WalletRepository.CreateWallet(wallet); err != nil {
		return err
	}

	return nil
}

func (i *implUser) GetUserWallet(ctx context.Context, userID string) (data []*dto.UserWalletResponseDTO, err error) {
	wallets, err := i.repository.WalletRepository.GetWalletsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, wallet := range wallets {
		currency, err := i.repository.CurrencyRepository.GetCurrencyByID(ctx, wallet.CurrencyID)
		if err != nil {
			return nil, err
		}

		data = append(data, &dto.UserWalletResponseDTO{
			UserWallet: dto.UserWalletDTO{
				WalletID:   wallet.ID,
				WalletName: wallet.WalletName,
			},
			Currency: dto.UserWalletCurrencyDTO{
				CurrencyID:   currency.ID,
				CurrencyName: currency.CurrencyName,
				CurrencyCode: currency.CurrencyCode,
			},
		})
	}

	return data, nil
}
