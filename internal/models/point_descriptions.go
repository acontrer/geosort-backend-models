package models

type PointDescriptions struct {
	Id          int    `gorm:"column:id;not null;" json:"id" form:"point_descriptions_id"`
	Description string `gorm:"column:description;not null;" json:"description" form:"point_descriptions_description"`
	PointsId    int    `gorm:"column:points_id;not null;" json:"points_id" form:"point_descriptions_points_id"`
}
