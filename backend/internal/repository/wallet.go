package repository

import (
	"database/sql"
	"expense_app/internal/entity"
	"expense_app/internal/util/logger"
)

type (
	WalletRepository interface {
		CreateWallet(*entity.Wallet) error
	}

	implWallet struct {
		logger logger.Logger
		db     *sql.DB
	}
)

func NewWalletRepository(logger logger.Logger, db *sql.DB) (WalletRepository, error) {
	return &implWallet{logger: logger, db: db}, nil
}

func (i *implWallet) CreateWallet(wallet *entity.Wallet) error {
	query := `INSERT INTO 
	wallets (id, user_id, currency_id, wallet_name, created_by)
	VALUES ($1, $2, $3, $4, $5)`

	args := []interface{}{
		wallet.ID,
		wallet.UserID,
		wallet.CurrencyID,
		wallet.WalletName,
		wallet.UserID,
	}

	row := i.db.QueryRow(query, args...)
	if row.Err() != nil {
		i.logger.Errorf("error creating wallet: %+v", row.Err().Error())
		return row.Err()
	}

	return nil
}
