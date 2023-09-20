package service

import (
	"expense_app/internal/controller/dto"
	"expense_app/internal/repository"
	"expense_app/internal/util/logger"
)

type (
	CurrencyService interface {
		GetCurrencies() ([]dto.CurrencyResponseDTO, error)
	}

	implCurrency struct {
		logger     logger.Logger
		repository repository.Holder
	}
)

func NewCurrencyService(logger logger.Logger, repository repository.Holder) (CurrencyService, error) {
	return &implCurrency{logger: logger, repository: repository}, nil
}

func (i *implCurrency) GetCurrencies() (result []dto.CurrencyResponseDTO, err error) {
	rows, err := i.repository.CurrencyRepository.GetCurrencies()
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		result = append(result, dto.CurrencyResponseDTO{
			ID:                 row.ID,
			CurrencyName:       row.CurrencyName,
			CurrencyCode:       row.CurrencyCode,
			CurrencyNumber:     row.CurrencyNumber,
			CurrencyMinorUnits: row.CurrencyMinorUnits,
		})
	}

	return result, nil
}
