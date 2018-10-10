package models

import (
	"github.com/jinzhu/gorm"

	"time"
)

type Vehicles struct {
	Id                  int          `gorm:"column:id;not null;" json:"id" form:"vehicles_id"`
	VehicleTypesId      int          `gorm:"column:vehicle_types_id;not null;" json:"vehicle_types_id" form:"vehicles_vehicle_types_id"`
	VehicleType         VehicleTypes `gorm:"foreignkey:VehicleTypesId;" json:"vehicle_type" form:"vehicles_vehicle_type"`
	EnterprisesId       int          `gorm:"column:enterprises_id;not null;" json:"enterprises_id" form:"vehicles_enterprises_id"`
	Enterprise          Enterprises  `gorm:"foreignkey:EnterprisesId;" json:"enterprise" form:"vehicles_enterprise"`
	Patent              string       `gorm:"column:patent;not null;" json:"patent" form:"vehicles_patent"`
	State               int          `gorm:"column:state;not null;" json:"state" form:"vehicles_state"`
	Certificate         time.Time    `gorm:"column:certificate;not null;" json:"certificate" form:"vehicles_certificate"`
	TechnicalReview     time.Time    `gorm:"column:technical_review;not null;" json:"technical_review" form:"vehicles_technical_review"`
	Insurance           time.Time    `gorm:"column:insurance;not null;" json:"insurance" form:"vehicles_insurance"`
	ObligatoryInsurance time.Time    `gorm:"column:obligatory_insurance;not null;" json:"obligatory_insurance" form:"vehicles_obligatory_insurance"`
}

func (v *Vehicles) GetCapacity() float64 {
	return v.VehicleType.Height * v.VehicleType.Length * v.VehicleType.Width
}

func (v *Vehicles) Expand(data *gorm.DB) error {
	if err := data.Model(v).Related(&v.VehicleType).Error; err != nil {
		return err
	}

	if err := data.Model(v).Related(&v.Enterprise).Error; err != nil {
		return err
	}

	return nil
}
