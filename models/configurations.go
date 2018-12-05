package models

import "time"

type Configurations struct {
	Id                   int       `gorm:"column:id;not null;" json:"id" form:"configurations_id"`
	PlanningDate         time.Time `gorm:"column:planning_date;not null;" json:"planning_date" form:"configurations_planning_date"`
	PlanningRadius       float64   `gorm:"column:planning_radius;not null;" json:"planning_radius" form:"configurations_planning_radius"`
	PlanningExpandRadius int       `gorm:"column:planning_expand_radius;not null;" json:"planning_expand_radius" form:"configurations_planning_expand_radius"`
	DeliveryTime         int       `gorm:"column:delivery_time;not null;" json:"delivery_time" form:"configurations_delivery_time"`
}
