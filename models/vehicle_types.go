package models

type VehicleTypes struct {
	Id     int     `gorm:"column:id;not null;" json:"id" form:"vehicle_types_id"`
	Name   string  `gorm:"column:name;not null;" json:"name" form:"vehicle_types_name"`
	Width  float64 `gorm:"column:width;not null;" json:"width" form:"vehicle_types_width"`
	Length float64 `gorm:"column:length;not null;" json:"length" form:"vehicle_types_length"`
	Height float64 `gorm:"column:height;not null;" json:"height" form:"vehicle_types_height"`
}
