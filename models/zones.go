package models

type Zones struct {
	Id       int        `gorm:"column:id;not null;" json:"id" form:"zones_id"`
	Zone     string     `gorm:"column:zone;not null;" json:"zone" form:"zones_zone"`
	Communes []Communes `gorm:"foreignkey:ZonesId;association_foreignkey:Id;" json:"communes"`
}
