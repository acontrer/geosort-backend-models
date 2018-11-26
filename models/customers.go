package models

import "github.com/jinzhu/gorm"

type Customers struct {
	Id        int    `gorm:"column:id;not null;" json:"id" form:"customers_id"`
	Rut       string `gorm:"column:rut;not null;" json:"rut" form:"customers_rut"`
	FirstName string `gorm:"column:first_name;not null;" json:"first_name" form:"customers_first_name"`
	LastName  string `gorm:"column:last_name;not null;" json:"last_name" form:"customers_last_name"`
}

func (c *Customers) GetCustomerByRut(data *gorm.DB) bool {
	if err := data.Where("rut = ?", c.Rut).Find(&c).Error; err != nil {
		return false
	} else {
		return true
	}
}
