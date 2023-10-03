package repository

import (
	"context"
	"database/sql"
	"expense_app/internal/entity"
	"expense_app/internal/util/logger"
)

type (
	WalletRepository interface {
		CreateWallet(*entity.Wallet) error
		GetWalletsByUserID(ctx context.Context, userID string) (wallets []entity.Wallet, err error)
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

func (i *implWallet) GetWalletsByUserID(ctx context.Context, userID string) (wallets []entity.Wallet, err error) {
	query := `SELECT id, wallet_name, currency_id FROM wallets WHERE user_id = $1 AND is_deleted=false`

	args := []interface{}{
		userID,
	}

	rows, err := i.db.QueryContext(ctx, query, args...)
	if err != nil {
		i.logger.Errorf("error getting wallet: %+v", err)
		return nil, err
	}

	if rows.Err() != nil {
		i.logger.Errorf("error getting wallet: %+v", err)
		return nil, rows.Err()
	}

	for rows.Next() {
		var wallet entity.Wallet
		err = rows.Scan(&wallet.ID, &wallet.WalletName, &wallet.CurrencyID)
		if err != nil {
			if err != nil {
				i.logger.Error("scan row into struct error")
				return nil, err
			}
		}

		wallets = append(wallets, wallet)
	}

	return wallets, nil
}
