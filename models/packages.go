package models

import (
	"github.com/jinzhu/gorm"
)

type Packages struct {
	Id             int64        `gorm:"column:id;not null;" json:"id" form:"packages_id"`
	Sku            int          `gorm:"column:sku;not null;" json:"sku" form:"packages_sku"`
	Description    string       `gorm:"column:description;not null;" json:"description" form:"packages_description"`
	Quantity       int          `gorm:"column:quantity;not null;" json:"quantity" form:"packages_quantity"`
	Width          float64      `gorm:"column:width;not null;" json:"width" form:"packages_width"`
	Length         float64      `gorm:"column:length;not null;" json:"length" form:"packages_length"`
	Height         float64      `gorm:"column:height;not null;" json:"height" form:"packages_height"`
	Weight         float64      `gorm:"column:weight;not null;" json:"weight" form:"packages_weight"`
	RotatyTypesId  int          `gorm:"column:rotaty_types_id;not null;" json:"rotaty_types_id" form:"packages_rotaty_types_id"`
	Stackable      int          `gorm:"column:stackable;not null;" json:"stackable" form:"packages_stackable"`
	Price          int          `gorm:"column:price;not null;" json:"price" form:"packages_price"`
	PackageTypesId int          `gorm:"column:package_types_id;not null;" json:"package_types_id" form:"packages_package_types_id"`
	PackageType    PackageTypes `gorm:"foreignkey:PackageTypesId;" json:"package_type" form:"packages_package_type"`
	SubordersId    int64        `gorm:"column:suborders_id;not null;" json:"suborders_id" form:"packages_suborders_id"`
	Priority       int          `gorm:"column:priority;not null;" json:"priority" form:"packages_priority"`
	LastState      *int         `gorm:"column:last_state;" json:"last_state" form:"packages_last_state"`
	LastStateModel States       `gorm:"foreignkey:LastState;" json:"last_state_model" form:"packages_last_state_model"`
	States         []States     `gorm:"foreignkey:PackagesId;association_foreignkey:Id" json:"state" form:"packages_states"`
}

func (p *Packages) Expand(data *gorm.DB) error {
	if err := data.Model(p).Related(&p.LastStateModel).Error; err != nil {
		return err
	} else {
		if err := p.LastStateModel.Expand(data); err != nil {
			return err
		}
	}

	if err := data.Model(p).Related(&p.States).Error; err != nil {
		return err
	} else {
		for i := range p.States {
			if err := p.States[i].Expand(data); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *Packages) GetVolume() float64 {
	return p.Height * p.Length * p.Width
}
