package models

type RestrictionAreas struct {
	Id      int    `gorm:"column:id;not null;" json:"id" form:"restriction_areas_id"`
	Name    string `gorm:"column:name;not null;" json:"name" form:"restriction_areas_name"`
	Polygon string `gorm:"column:polygon;not null;" json:"polygon" form:"restriction_areas_polygon"`
}
