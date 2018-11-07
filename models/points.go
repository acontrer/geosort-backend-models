package models

import (
	"github.com/dwladdimiroc/geosort-backend-models/utils"
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
			return utils.NewError(err, "geo details")
		}
	} else {
		for i := range p.GeoDetails {
			if err := p.GeoDetails[i].Expand(data); err != nil {
				return utils.NewError(err, "geo details expand")
			}
		}
	}

	if err := data.Model(p).Related(&p.Street).Error; err != nil {
		return utils.NewError(err, "street")
	} else {
		if err := p.Street.Expand(data); err != nil {
			return utils.NewError(err, "street expand")
		}
	}

	return nil
}

func (p *Points) GetPointByStreet(db *gorm.DB, streetName, streetNumber string) bool {
	var street Streets
	if err := db.Where("name = ?", streetName).Find(&street).Error; err != nil {
		return false
	} else {
		if err := db.Where("streets_id = ? AND street_number = ?", street.Id, streetNumber).Find(&p).Error; err != nil {
			return false
		} else {
			return true
		}
	}
}

//TODO: Where is unnecessary. Make the condition in join clause.
func (p *Points) CheckRestrictionAreas(data *gorm.DB) ([]RestrictionAreas, error) {
	var ra []RestrictionAreas
	tx := data.Joins("JOIN points ON (st_within(st_setsrid(st_makepoint(points.longitude, points.latitude), 4267), restriction_areas.polygon::geometry))")
	tx = tx.Where(p)
	err := tx.Find(&ra).Error
	return ra, err
}
