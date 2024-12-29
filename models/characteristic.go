package models

import "github.com/google/uuid"

type Characteristic struct {
	Meta  Meta      `json:"meta"`
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Value string    `json:"value"`
}
