package models

import (
	"github.com/dwladdimiroc/geosort-backend-models/utils"
	"github.com/jinzhu/gorm"
)

type Communes struct {
	Id        int     `gorm:"column:id;not null;" json:"id" form:"communes_id"`
	Name      string  `gorm:"column:name;not null;" json:"name" form:"communes_name"`
	ZonesId   int     `gorm:"column:zones_id;not null;" json:"zones_id" form:"communes_zones_id"`
	RegionsId int     `gorm:"column:regions_id;not null;" json:"regions_id" form:"communes_regions_id"`
	Region    Regions `gorm:"foreignkey:RegionsId" json:"region" form:"communes_region"`
}

func (c *Communes) Expand(data *gorm.DB) error {
	if err := data.Model(c).Related(&c.Region).Error; err != nil {
		return utils.NewError(err, "region")
	} else {
		if err := c.Region.Expand(data); err != nil {
			return utils.NewError(err, "region expand")
		}
	}

	return nil
}
