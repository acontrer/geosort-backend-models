package models

type Restrictions struct {
	Id                 int `gorm:"column:id;not null;" json:"id" form:"restrictions_id"`
	TimeWindowsId      int `gorm:"column:time_windows_id;not null;" json:"time_windows_id" form:"restrictions_time_windows_id"`
	AreaRestrictionsId int `gorm:"column:area_restrictions_id;not null;" json:"area_restrictions_id" form:"restrictions_area_restrictions_id"`
}
