package models

import "github.com/jinzhu/gorm"

type StateSubtypes struct {
	Id           int        `gorm:"column:id;not null;" json:"id" form:"state_subtypes_id"`
	Subtype      string     `gorm:"column:subtype;not null;" json:"subtype" form:"state_subtypes_subtype"`
	StateTypesId int        `gorm:"column:state_types_id;not null;" json:"state_types_id" form:"state_types_state_subtypes_id"`
	Code         string     `gorm:"column:code;not null;" json:"code" form:"state_subtypes_code"`
	Subcode      string     `gorm:"column:subcode;not null;" json:"subcode" form:"state_types_subcode"`
	StateTypes   StateTypes `gorm:"foreignkey:StateSubtypesId" json:"state_types" form:"state_subtypes_state_type"`
}

func (st *StateSubtypes) Expand(data *gorm.DB) error {
	if err := data.Model(st).Related(&st.StateTypes).Error; err != nil {
		return err
	}
	return nil
}
