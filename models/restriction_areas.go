package models

import "github.com/jinzhu/gorm"

type RestrictionAreas struct {
	Id      int    `gorm:"column:id;not null;" json:"id" form:"restriction_areas_id"`
	Name    string `gorm:"column:name;not null;" json:"name" form:"restriction_areas_name"`
	Polygon string `gorm:"column:polygon;not null;" json:"polygon" form:"restriction_areas_polygon"`
}

func (r *RestrictionAreas) Add(data *gorm.DB) (error) {
	if err := data.Raw("INSERT INTO restriction_areas (name, polygon) VALUES (?, st_geomfromtext(?, 4267)) returning restriction_areas.id", r.Name, r.Polygon).Scan(&r).Error; err != nil {
		return err
	} else {
		return nil
	}
}
