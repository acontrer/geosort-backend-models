package models

type StateSubtypes struct {
	Id      int    `gorm:"column:id;not null;" json:"id" form:"state_subtypes_id"`
	Subtype string `gorm:"column:subtype;not null;" json:"subtype" form:"state_subtypes_subtype"`
	Code    string `gorm:"column:code;not null;" json:"code" form:"state_subtypes_code"`
}
