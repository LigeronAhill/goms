package models

type Price struct {
	Value     float64     `json:"value"`
	Currency  MetaWrapper `json:"currency"`
	PriceType PriceType   `json:"priceType,omitempty"`
}
