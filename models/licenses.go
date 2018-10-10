package models

import (
	"time"
)

type Licenses struct {
	Id        int       `gorm:"column:id;not null;" json:"id" form:"licenses_id"`
	Type      string    `gorm:"column:type;not null;" json:"type" form:"licenses_type"`
	Expire    time.Time `gorm:"column:expire;not null;" json:"expire" form:"licenses_expire"`
	DriversId int       `gorm:"column:drivers_id;not null;" json:"drivers_id" form:"licenses_drivers_id"`
}
