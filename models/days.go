package models

type Days struct {
	Id      int    `gorm:"column:id;not null;" json:"id" form:"days_id"`
	DayName string `gorm:"column:day_name;not null;" json:"day_name" form:"days_day_name"`
}
