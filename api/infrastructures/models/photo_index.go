package models

import "fmt"

type PhotoIndex struct {
	PhotoID          string `json:"photo_id"`
	Name             string `json:"name"`
	ImportedAt       int64  `json:"imported_at"`
	DateTimeOriginal int64  `json:"date_time_original"`
	DescriptionJa    string `json:"description_ja"`
	DescriptionEn    string `json:"description_en"`
}

func (m PhotoIndex) IndexName() string {
	return "photo"
}

func (m PhotoIndex) DocumentID() string {
	return fmt.Sprintf("%s", m.PhotoID)
}
