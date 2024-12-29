package models

import "github.com/google/uuid"

type Product struct {
	AccountID           uuid.UUID   `json:"account_id"`
	Archived            bool        `json:"archived"`
	Article             string      `json:"article,omitempty"`
	Attributes          []Attribute `json:"attributes,omitempty"`
	Barcodes            []Barcode   `json:"barcodes,omitempty"`
	BuyPrice            Price       `json:"buyPrice,omitempty"`
	Code                string      `json:"code,omitempty"`
	Country             MetaWrapper `json:"country,omitempty"`
	Description         string      `json:"description,omitempty"`
	DiscountProhibited  bool        `json:"discountProhibited,omitempty"`
	EffectiveVat        int         `json:"effectiveVat,omitempty"`
	EffectiveVatEnabled bool        `json:"effectiveVatEnabled,omitempty"`
	ExternalCode        string      `json:"externalCode,omitempty"`
	Files               MetaWrapper `json:"files,omitempty"`
	Group               MetaWrapper `json:"group,omitempty"`
	ID                  uuid.UUID   `json:"id"`
	Images              MetaWrapper `json:"images,omitempty"`
	IsSerialTrackable   bool        `json:"isSerialTrackable,omitempty"`
	Meta                Meta        `json:"meta"`
	MinPrice            Price       `json:"minPrice,omitempty"`
	MinimumBalance      int         `json:"minimumBalance,omitempty"`
	Name                string      `json:"name"`
	Owner               MetaWrapper `json:"owner,omitempty"`
	PathName            string      `json:"pathName"`
	PaymentItemType     string      `json:"paymentItemType,omitempty"`
	ProductFolder       MetaWrapper `json:"productFolder,omitempty"`
	SalePrices          []Price     `json:"salePrices,omitempty"`
	Shared              bool        `json:"shared"`
	Supplier            MetaWrapper `json:"supplier,omitempty"`
	Uom                 MetaWrapper `json:"uom,omitempty"`
	Updated             MSDate      `json:"updated"`
	UseParentVat        bool        `json:"useParentVat,omitempty"`
	VariantsCount       int         `json:"variantsCount"`
	Vat                 int         `json:"vat,omitempty"`
	VatEnabled          bool        `json:"vatEnabled,omitempty"`
	Volume              float64     `json:"volume,omitempty"`
	Weight              float64     `json:"weight,omitempty"`
}
