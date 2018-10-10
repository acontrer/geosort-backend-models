package models

type PackageTypes struct {
	Id   int    `gorm:"column:id;not null;" json:"id" form:"package_types_id"`
	Type string `gorm:"column:type;not null;" json:"type" form:"package_types_type"`
}
