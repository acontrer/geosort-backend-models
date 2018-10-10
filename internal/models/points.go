package models

import (
	"github.com/jinzhu/gorm"
)

type Points struct {
	Id           int          `gorm:"column:id;not null;" json:"id" form:"points_id"`
	GeoDetails   []GeoDetails `gorm:"foreignkey:PointsId;association_foreignkey:Id;" json:"geo_details" form:"points_geo_details"`
	Latitude     float64      `gorm:"column:latitude;not null;" json:"latitude" form:"points_latitude"`
	Longitude    float64      `gorm:"column:longitude;not null;" json:"longitude" form:"points_longitude"`
	StreetsId    int          `gorm:"column:streets_id;not null;" json:"streets_id" form:"points_streets_id"`
	Street       Streets      `gorm:"foreignkey:StreetsId" json:"street" form:"points_street"`
	StreetNumber string       `gorm:"column:street_number;not null;" json:"street_number" form:"points_street_number"`
}

func (p *Points) Expand(data *gorm.DB) error {
	if err := data.Model(p).Related(&p.GeoDetails).Error; err != nil {
		if err.Error() != "record not found" {
			return err
		}
	} else {
		for i := range p.GeoDetails {
			if err := p.GeoDetails[i].Expand(data); err != nil {
				return err
			}
		}
	}

	if err := data.Model(p).Related(&p.Street).Error; err != nil {
		return err
	} else {
		if err := p.Street.Expand(data); err != nil {
			return err
		}
	}

	return nil
}
