package models

import (
	"github.com/google/uuid"
)

type Attribute struct {
	Meta  Meta      `json:"meta"`
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Type  string    `json:"type"`
	Value any       `json:"value"`
}

func (a Attribute) GetValue() string {
	switch v := a.Value.(type) {
	case string:
		return v
	case CustomValue:
		return v.Name
	default:
		return ""
	}
}

type CustomValue struct {
	Meta Meta   `json:"meta"`
	Name string `json:"name"`
}
