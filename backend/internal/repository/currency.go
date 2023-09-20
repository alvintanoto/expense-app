package repository

import (
	"context"
	"database/sql"
	"errors"
	"expense_app/internal/entity"
	"expense_app/internal/util/logger"
)

type (
	CurrencyRepository interface {
		GetCurrencies() (currencies []entity.Currency, err error)
	}

	implCurrency struct {
		logger logger.Logger
		db     *sql.DB
	}
)

func NewCurrencyRepository(logger logger.Logger, db *sql.DB) (CurrencyRepository, error) {
	return &implCurrency{logger: logger, db: db}, nil
}

func (i *implCurrency) GetCurrencies() (currencies []entity.Currency, err error) {
	query := `SELECT id, currency_name, currency_code, currency_number, currency_minor_units FROM currencies ORDER BY currency_name ASC`

	rows, err := i.db.QueryContext(context.Background(), query)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			i.logger.Errorf("no record found: %v", err)
			return nil, ErrorRecordNotFound
		default:
			i.logger.Errorf("scan row into struct error: %v", err)
			return nil, err
		}
	}

	for rows.Next() {
		var currency entity.Currency

		err = rows.Scan(&currency.ID, &currency.CurrencyName, &currency.CurrencyCode, &currency.CurrencyName, &currency.CurrencyMinorUnits)
		if err != nil {
			if err != nil {
				i.logger.Errorf("scan row into struct error: %v", err)
				return nil, err
			}
		}

		currencies = append(currencies, currency)
	}

	return currencies, nil
}
