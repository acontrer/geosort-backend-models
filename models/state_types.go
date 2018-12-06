package models

import "github.com/jinzhu/gorm"

type StateTypes struct {
	Id           int             `gorm:"column:id;not null;" json:"id" form:"state_types_id"`
	Type         string          `gorm:"column:type;not null;" json:"type" form:"state_types_type"`
	Code         string          `gorm:"column:code;not null;" json:"code" form:"state_types_code"`
	StateSubtype []StateSubtypes `gorm:"foreignkey:StateSubtypesId;association_autoupdate:false" json:"state_subtype" form:"state_types_state_subtype"`
}

func (st *StateTypes) Expand(data *gorm.DB) error {
	if err := data.Model(st).Related(&st.StateSubtype).Error; err != nil {
		return err
	}
	return nil
}
