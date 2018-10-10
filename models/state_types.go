package models

type StateTypes struct {
	Id       int    `gorm:"column:id;not null;" json:"id" form:"state_types_id"`
	Type     string `gorm:"column:type;not null;" json:"type" form:"state_types_type"`
	SubState string `gorm:"column:sub_state;not null;" json:"sub_state" form:"state_types_sub_state"`
}
