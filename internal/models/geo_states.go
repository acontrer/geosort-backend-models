package models

type GeoStates struct {
	Id    int    `gorm:"column:id;not null;" json:"id" form:"geo_states_id"`
	State string `gorm:"column:state;not null;" json:"state" form:"geo_states_state"`
}
