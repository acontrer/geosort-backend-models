package models

type Customers struct {
	Id        int    `gorm:"column:id;not null;" json:"id" form:"customers_id"`
	Rut       string `gorm:"column:rut;not null;" json:"rut" form:"customers_rut"`
	FirstName string `gorm:"column:first_name;not null;" json:"first_name" form:"customers_first_name"`
	LastName  string `gorm:"column:last_name;not null;" json:"last_name" form:"customers_last_name"`
}
