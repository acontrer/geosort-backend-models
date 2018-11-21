package models

import "github.com/jinzhu/gorm"

type Stores struct {
	Id               int            `gorm:"column:id;not null;" json:"id" form:"stores_id"`
	Name             string         `gorm:"column:name;not null;" json:"name" form:"stores_name"`
	DeliveryPointsId int            `gorm:"column:delivery_points_id;not null;" json:"delivery_points_id" form:"stores_delivery_points_id"`
	DeliveryPoint    DeliveryPoints `gorm:"foreignkey:DeliveryPointsId;" json:"delivery_point" form:"stores_delivery_point"`
	PointsId         int            `gorm:"column:points_id;not null;" json:"points_id" form:"stores_points_id"`
	Point            Points         `gorm:"foreignkey:PointsId;" json:"point" form:"stores_point"`
}

func (s *Stores) Expand(data *gorm.DB) error {
	if err := data.Model(s).Related(&s.Point).Error; err != nil {
		return err
	}

	if err := data.Model(s).Related(&s.DeliveryPoint).Error; err != nil {
		return err
	}

	return nil
}
