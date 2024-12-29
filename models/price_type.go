package models

type PriceType struct {
	Meta         Meta   `json:"meta"`
	ID           string `json:"id"`
	Name         string `json:"name"`
	ExternalCode string `json:"externalCode"`
}
