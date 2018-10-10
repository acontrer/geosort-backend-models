package models

type RotatyTypes struct {
	Id          int    `gorm:"column:id;not null;" json:"id" form:"rotaty_types_id"`
	Description string `gorm:"column:description;not null;" json:"description" form:"rotaty_types_description"`
}
