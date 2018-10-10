package models

type Preferences struct {
	Id        int `gorm:"column:id;not null;" json:"id" form:"preferences_id"`
	DriversId int `gorm:"column:drivers_id;not null;" json:"drivers_id" form:"preferences_drivers_id"`
	ZonesId   int `gorm:"column:zones_id;not null;" json:"zones_id" form:"preferences_zones_id"`
}
