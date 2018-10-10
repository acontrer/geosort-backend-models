package models

type RestrictionAreasDays struct {
	Id                 int `gorm:"column:id;not null;" json:"id" form:"restriction_areas_days_id"`
	RestrictionAreasId int `gorm:"column:restriction_areas_id;not null;" json:"restriction_areas_id" form:"restriction_areas_days_restriction_areas_id"`
	DaysId             int `gorm:"column:days_id;not null;" json:"days_id" form:"restriction_areas_days_days_id"`
}
