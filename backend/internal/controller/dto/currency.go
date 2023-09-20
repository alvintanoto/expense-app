package dto

type (
	CurrencyResponseDTO struct {
		ID                 string `json:"id"`
		CurrencyName       string `json:"currency_name"`
		CurrencyCode       string `json:"currency_code"`
		CurrencyNumber     string `json:"currency_number"`
		CurrencyMinorUnits int    `json:"currency_minor_units"`
	}
)
