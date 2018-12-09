package models

import (
	"github.com/jinzhu/gorm"

	"time"
)

type States struct {
	Id           int        `gorm:"column:id;not null;" json:"id" form:"states_id"`
	StateTypesId int        `gorm:"column:state_types_id;not null;" json:"state_types_id" form:"states_state_types_id"`
	StateType    StateTypes `gorm:"foreignkey:StateTypesId;association_autoupdate:false" json:"state_type" form:"states_state_type"`
	Description  string     `gorm:"column:description;not null;" json:"description" form:"states_description"`
	CreatedAt    time.Time  `gorm:"column:created_at;not null;" json:"created_at" form:"states_created_at"`
	PackagesId   int64      `gorm:"column:packages_id;not null;" json:"packages_id" form:"states_packages_id"`
	MatCode      string     `gorm:"column:mat_code;not null;" json:"mat_code" form:"states_mat_code"`
}

func (s *States) Expand(data *gorm.DB) error {
	if err := data.Model(s).Related(&s.StateType).Error; err != nil {
		return err
	} else {
		if err := s.StateType.Expand(data); err != nil {
			return err
		}
	}
	return nil
}
