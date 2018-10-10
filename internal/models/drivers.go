package models

import (
	"github.com/jinzhu/gorm"
)

type Drivers struct {
	Id            int      `gorm:"column:id;not null;" json:"id" form:"drivers_id"`
	Rut           string   `gorm:"column:rut;not null;" json:"rut" form:"drivers_rut"`
	FirstName     string   `gorm:"column:first_name;not null;" json:"first_name" form:"drivers_first_name"`
	LastName      string   `gorm:"column:last_name;not null;" json:"last_name" form:"drivers_last_name"`
	Phone         string   `gorm:"column:phone;not null;" json:"phone" form:"drivers_phone"`
	AccountsId    int      `gorm:"column:accounts_id;not null;" json:"accounts_id" form:"drivers_accounts_id"`
	Account       Accounts `gorm:"foreignkey:AccountsId;" json:"account" form:"account"`
	EnterprisesId int      `gorm:"column:enterprises_id;not null;" json:"enterprises_id" form:"drivers_enterprises_id"`
	Zones         []Zones  `gorm:"many2many:preferences;" json:"preferences" form:"zones"`
}

func (d *Drivers) GetZonesId() []int {
	zonesId := make([]int, len(d.Zones))
	for i := range d.Zones {
		zonesId[i] = d.Zones[i].Id
	}
	return zonesId
}

func (d *Drivers) Expand(data *gorm.DB) error {
	if err := data.Model(d).Related(&d.Account).Error; err != nil {
		return err
	} else {
		if err := d.Account.Expand(data); err != nil {
			return err
		}
	}

	return nil
}
