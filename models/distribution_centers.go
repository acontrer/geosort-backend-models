package models

import (
	"time"

	"strings"

	"github.com/dwladdimiroc/geosort-backend-models/utils"
	"github.com/jinzhu/gorm"
)

type DistributionCenters struct {
	Id               int       `gorm:"column:id;not null;" json:"id" form:"distribution_centers_id"`
	Name             string    `gorm:"column:name;not null;" json:"name" form:"distribution_centers_name"`
	StartWorkingDay  time.Time `gorm:"column:start_working_day;not null;" json:"start_working_day" form:"distribution_centers_start_working_day"`
	FinishWorkingDay time.Time `gorm:"column:finish_working_day;not null;" json:"finish_working_day" form:"distribution_centers_finish_working_day"`
	PointsId         int       `gorm:"column:points_id;not null;" json:"points_id" form:"distribution_centers_points_id"`
	Point            Points    `gorm:"foreignkey:PointsId;" json:"point" form:"distribution_centers_point"`
}

func (dc *DistributionCenters) GetDistribuionCenterByName(data *gorm.DB, name string) bool {
	name = strings.ToLower(utils.ValidString(name))
	if err := data.Where("lower(name) = ?", name).Find(&dc).Error; err != nil {
		return false
	} else {
		return true
	}
}

func (dc *DistributionCenters) Expand(data *gorm.DB) error {
	if err := data.Model(dc).Related(&dc.Point).Error; err != nil {
		return err
	} else {
		if err := dc.Point.Expand(data); err != nil {
			return err
		}
	}

	return nil
}
