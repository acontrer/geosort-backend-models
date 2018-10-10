package models

import (
	"github.com/jinzhu/gorm"
)

type Regions struct {
	Id          int       `gorm:"column:id;not null;" json:"id" form:"regions_id"`
	Name        string    `gorm:"column:name;not null;" json:"name" form:"regions_name"`
	CountriesId int       `gorm:"column:countries_id;not null;" json:"countries_id" form:"regions_countries_id"`
	Country     Countries `gorm:"foreignkey:CountriesId;" json:"country" form:"regions_country"`
}

func (r *Regions) Expand(data *gorm.DB) error {
	if err := data.Model(r).Related(&r.Country).Error; err != nil {
		return err
	} else {
		return nil
	}
}
