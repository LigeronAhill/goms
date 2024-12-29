package models

import "github.com/google/uuid"

// ProductFolder - Группа товаров
//
// Средствами JSON API можно создавать и обновлять сведения о Группах товаров, запрашивать списки Групп товаров и сведения по отдельным Группам товаров. Кодом сущности для Группы товаров в составе JSON API является ключевое слово productfolder. По данной сущности можно осуществлять контекстный поиск с помощью специального параметра search. Подробнее можно узнать по ссылке.
//
// Поиск среди объектов групп товаров на соответствие поисковой строке будет осуществлен по следующим полям:
//
// - по наименованию Группы товаров (name)
//
// - по коду Группы товаров (code)
type ProductFolder struct {
	AccountId           uuid.UUID   `json:"accountId,omitempty"`
	Archived            bool        `json:"archived,omitempty"`
	Code                string      `json:"code,omitempty"`
	Description         string      `json:"description,omitempty"`
	EffectiveVat        int         `json:"effectiveVat,omitempty"`
	EffectiveVatEnabled bool        `json:"effectiveVatEnabled,omitempty"`
	ExternalCode        string      `json:"externalCode,omitempty"`
	Group               MetaWrapper `json:"group,omitempty"`
	ID                  uuid.UUID   `json:"id,omitempty"`
	Meta                Meta        `json:"meta,omitempty"`
	Name                string      `json:"name"`
	Owner               MetaWrapper `json:"owner,omitempty"`
	PathName            string      `json:"pathName,omitempty"`
	ProductFolder       MetaWrapper `json:"productFolder,omitempty"`
	Shared              bool        `json:"shared,omitempty"`
	TaxSystem           string      `json:"taxSystem,omitempty"`
	Updated             string      `json:"updated,omitempty"`
	UseParentVat        bool        `json:"useParentVat,omitempty"`
	Vat                 int         `json:"vat,omitempty"`
	VatEnabled          bool        `json:"vatEnabled,omitempty"`
}
