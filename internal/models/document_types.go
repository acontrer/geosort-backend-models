package models

type DocumentTypes struct {
	Id   int    `gorm:"column:id;not null;" json:"id" form:"document_types_id"`
	Type string `gorm:"column:type;not null;" json:"type" form:"document_types_type"`
}
