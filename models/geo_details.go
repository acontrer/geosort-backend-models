package models

import (
	"github.com/jinzhu/gorm"

	"time"
)

type GeoDetails struct {
	Id          int       `gorm:"column:id;not null;" json:"id" form:"geo_details_id"`
	Timestamp   time.Time `gorm:"column:timestamp;not null;" json:"timestamp" form:"geo_details_timestamp"`
	PointsId    int       `gorm:"column:points_id;not null;" json:"points_id" form:"geo_details_points_id"`
	GeoStatesId int       `gorm:"column:geo_states_id;not null;" json:"geo_states_id" form:"geo_details_geo_states_id"`
	GeoState    GeoStates `gorm:"foreignkey:GeoStateId;" json:"geo_state" form:"geo_details_geo_state"`
	AccountsId  int       `gorm:"column:accounts_id;not null;" json:"accounts_id" form:"geo_details_accounts_id"`
	Account     Accounts  `gorm:"foreignkey:AccountsId;" json:"account" form:"geo_details_account"`
}

func (gd *GeoDetails) Expand(data *gorm.DB) error {
	if err := data.Model(gd).Related(&gd.GeoState).Error; err != nil {
		return err
	}

	if err := data.Model(gd).Related(&gd.Account).Error; err != nil {
		return err
	} else {
		if err := gd.Account.Expand(data); err != nil {
			return err
		}
	}

	return nil
}
