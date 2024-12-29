package models

type Meta struct {
	Href         string `json:"href"`
	MetadataHref string `json:"metadataHref,omitempty"`
	Type         string `json:"type,omitempty"`
	MediaType    string `json:"mediaType"`
	UuidHref     string `json:"uuidHref,omitempty"`
	DownloadHref string `json:"downloadHref,omitempty"`
	Size         int    `json:"size,omitempty"`
	Limit        int    `json:"limit,omitempty"`
	Offset       int    `json:"offset,omitempty"`
}

type MetaWrapper struct {
	Meta Meta `json:"meta"`
}
