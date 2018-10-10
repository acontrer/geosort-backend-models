package models

type AccountTypes struct {
	Id   int    `gorm:"column:id;not null;" json:"id" form:"account_types_id"`
	Type string `gorm:"column:type;not null;" json:"type" form:"account_types_type"`
}
