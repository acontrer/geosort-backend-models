package models

type Enterprises struct {
	Id   int    `gorm:"column:id;not null;" json:"id" form:"enterprises_id"`
	Rut  string `gorm:"column:rut;not null;" json:"rut" form:"enterprises_rut"`
	Name string `gorm:"column:name;not null;" json:"name" form:"enterprises_name"`
}
