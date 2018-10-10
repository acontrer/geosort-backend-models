package models

type TravelTypes struct {
	Id   int    `gorm:"column:id;not null;" json:"id" form:"travel_types_id"`
	Type string `gorm:"column:type;not null;" json:"type" form:"travel_types_type"`
}
