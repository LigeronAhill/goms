package models

import (
	"github.com/google/uuid"
	"strings"
	"time"
)

type Uom struct {
	AccountID    uuid.UUID   `json:"account_id"`
	Code         string      `json:"code,omitempty"`
	Description  string      `json:"description,omitempty"`
	ExternalCode string      `json:"externalCode"`
	Group        MetaWrapper `json:"group,omitempty"`
	ID           uuid.UUID   `json:"id"`
	Meta         Meta        `json:"meta"`
	Name         string      `json:"name"`
	Owner        MetaWrapper `json:"owner,omitempty"`
	Shared       bool        `json:"shared"`
	Updated      MSDate      `json:"updated"`
}

func (u Uom) Endpoint() string {
	return "/entity/uom"
}

type MSDate struct {
	time.Time
}

func (c *MSDate) UnmarshalJSON(b []byte) (err error) {
	layout := "2006-01-02 15:04:05"

	s := strings.Trim(string(b), "\"") // remove quotes
	if s == "null" {
		return
	}
	c.Time, err = time.Parse(layout, s)
	return
}
