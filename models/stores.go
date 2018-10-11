package models

import "github.com/jinzhu/gorm"

type Stores struct {
	Id       int    `gorm:"column:id;not null;" json:"id" form:"stores_id"`
	Name     string `gorm:"column:name;not null;" json:"name" form:"stores_name"`
	PointsId int    `gorm:"column:points_id;not null;" json:"points_id" form:"stores_points_id"`
	Point    Points `gorm:"foreignkey:PointsId;" json:"point" form:"stores_point"`
}

func (s *Stores) Expand(data *gorm.DB) error {
	if err := data.Model(s).Related(&s.Point).Error; err != nil {
		return err
	} else {
		return nil
	}
}
