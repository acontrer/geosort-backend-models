package models

type TimeDetails struct {
	Id               int   `gorm:"column:id;not null;" json:"id" form:"time_details_id"`
	DeliveryPointsId int64 `gorm:"column:delivery_points_id;not null;" json:"delivery_points_id" form:"time_details_delivery_points_id"`
	TimeWindowsId    int   `gorm:"column:time_windows_id;not null;" json:"time_windows_id" form:"time_details_time_windows_id"`
}
