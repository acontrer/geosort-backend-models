package models

type Countries struct {
	Id   int    `gorm:"column:id;not null;" json:"id" form:"countries_id"`
	Name string `gorm:"column:name;not null;" json:"name" form:"countries_name"`
}
