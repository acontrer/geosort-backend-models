package models

type DeliveryMethods struct {
	Id     int    `gorm:"column:id;not null;" json:"id" form:"delivery_methods_id"`
	Method string `gorm:"column:method;not null;" json:"method" form:"delivery_methods_method"`
}
