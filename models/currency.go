package models

import "github.com/google/uuid"

type Currency struct {
	Archived       bool      `json:"archived"`
	Code           string    `json:"code"`
	Default        bool      `json:"default"`
	FullName       string    `json:"full_name,omitempty"`
	ID             uuid.UUID `json:"id"`
	Indirect       bool      `json:"indirect"`
	IsoCode        string    `json:"iso_code"`
	MajorUnit      Unit      `json:"major_unit,omitempty"`
	Margin         float64   `json:"margin,omitempty"`
	Meta           Meta      `json:"meta"`
	MinorUnit      Unit      `json:"minor_unit,omitempty"`
	Multiplicity   int       `json:"multiplicity"`
	Name           string    `json:"name"`
	Rate           float64   `json:"rate"`
	RateUpdateType string    `json:"rate_update_type"`
	System         bool      `json:"system"`
}

type Unit struct {
	Gender string `json:"gender,omitempty"`
	S1     string `json:"s1,omitempty"`
	S2     string `json:"s2,omitempty"`
	S3     string `json:"s3,omitempty"`
}
