package models

import (
	"github.com/jinzhu/gorm"

	"time"
)

type PlanningStatuses struct {
	Id                    int                 `gorm:"column:id;not null;" json:"id" form:"planning_statuses_id"`
	CreatedAt             time.Time           `gorm:"column:created_at;not null;" json:"created_at" form:"planning_statuses_state_types_id"`
	Status                string              `gorm:"column:status;not null;" json:"status" form:"planning_statuses_status"`
	Code                  int                 `gorm:"column:code;not null;" json:"code" form:"planning_statuses_code"`
	DistributionCentersId int                 `gorm:"column:distribution_centers_id;not null;" json:"distribution_centers_id" form:"planning_statuses_distribution_centers_id"`
	DistributionCenter    DistributionCenters `gorm:"foreignkey:DistributionCentersId" json:"distribution_center" form:"planning_statuses_distribution_center"`
	TravelTypesId         int                 `gorm:"column:travel_types_id;not null;" json:"travel_types_id" form:"planning_statuses_travel_types_id"`
	TravelType            TravelTypes         `gorm:"foreignkey:TravelTypesId" json:"travel_type" form:"planning_statuses_travel_type"`
}

func (sp *PlanningStatuses) Expand(data *gorm.DB) error {
	if err := data.Model(sp).Related(&sp.DistributionCenter).Error; err != nil {
		return err
	}

	if err := data.Model(sp).Related(&sp.TravelType).Error; err != nil {
		return err
	}

	return nil
}
