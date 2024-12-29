package models

import "github.com/google/uuid"

type Country struct {
	AccountID    uuid.UUID   `json:"account_id,omitempty"`
	Code         string      `json:"code,omitempty"`
	Description  string      `json:"description,omitempty"`
	ExternalCode string      `json:"external_code,omitempty"`
	Group        MetaWrapper `json:"group,omitempty"`
	ID           uuid.UUID   `json:"id"`
	Meta         Meta        `json:"meta"`
	Name         string      `json:"name"`
	Owner        MetaWrapper `json:"owner,omitempty"`
	Shared       bool        `json:"shared,omitempty"`
	Updated      MSDate      `json:"updated,omitempty"`
}
