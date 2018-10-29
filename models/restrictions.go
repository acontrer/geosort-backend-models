package models

import "github.com/jinzhu/gorm"

type Restrictions struct {
	Id                 int         `gorm:"column:id;not null;" json:"id" form:"restrictions_id"`
	TimeWindowsId      int         `gorm:"column:time_windows_id;not null;" json:"time_windows_id" form:"restrictions_time_windows_id"`
	AreaRestrictionsId int         `gorm:"column:area_restrictions_id;not null;" json:"area_restrictions_id" form:"restrictions_area_restrictions_id"`
	DaysId             int         `gorm:"column:days_id;not null;" json:"days_id" form:"restrictions_days_id"`
	TimeWindows        TimeWindows `gorm:"foreignkey:TimeWindowsId;" json:"time_windows" form:"restrictions_time_windows"`
	Days               Days        `gorm:"foreignkey:DaysId;" json:"days" form:"restrictions_days"`
}

func (r *Restrictions) Expand(data *gorm.DB) (error) {
	if err := data.Model(r).Related(&r.TimeWindows, "TimeWindows").Error; err != nil {
		return err
	}
	if err := data.Model(r).Related(&r.Days, "Days").Error; err != nil {
		return err
	}
	return nil
}
