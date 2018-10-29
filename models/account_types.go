package models

var (
	AdminType  = 1
	UserType   = 2
	DriverType = 3
	AllType    = []int{1, 2, 3}
)

type AccountTypes struct {
	Id   int    `gorm:"column:id;not null;" json:"id" form:"account_types_id"`
	Type string `gorm:"column:type;not null;" json:"type" form:"account_types_type"`
}
