package models

import "github.com/google/uuid"

type Variant struct {
	AccountID          uuid.UUID        `json:"account_id"`
	Archived           bool             `json:"archived"`
	Barcodes           []Barcode        `json:"barcodes,omitempty"`
	BuyPrice           Price            `json:"buyPrice,omitempty"`
	Characteristics    []Characteristic `json:"characteristics"`
	Code               string           `json:"code,omitempty"`
	Description        string           `json:"description,omitempty"`
	DiscountProhibited bool             `json:"discountProhibited,omitempty"`
	ExternalCode       string           `json:"externalCode,omitempty"`
	ID                 uuid.UUID        `json:"id"`
	Images             MetaWrapper      `json:"images,omitempty"`
	Meta               Meta             `json:"meta"`
	MinPrice           Price            `json:"minPrice,omitempty"`
	Name               string           `json:"name"`
	Product            MetaWrapper      `json:"product"`
	SalePrices         []Price          `json:"salePrices,omitempty"`
	Updated            MSDate           `json:"updated"`
}
