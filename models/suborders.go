package models

import (
	"github.com/dwladdimiroc/geosort-backend-models/utils"
	"github.com/jinzhu/gorm"
)

type Suborders struct {
	Id                int64           `gorm:"column:id;not null;" json:"id" form:"suborders_id"`
	DeliveryPointsId  int64           `gorm:"column:delivery_points_id;not null;" json:"delivery_points_id" form:"suborders_delivery_points_id"`
	DeliveryPoints    DeliveryPoints  `gorm:"foreignkey:DeliveryPointsId;" json:"delivery_points" form:"suborders_delivery_points"`
	Code              int64           `gorm:"column:code;not null;" json:"code" form:"suborders_code"`
	Do                int64           `gorm:"column:do;not null;" json:"do" form:"suborders_do"`
	DocumentTypesId   int             `gorm:"column:document_types_id;not null;" json:"document_types_id" form:"suborders_document_types_id"`
	DocumentType      DocumentTypes   `gorm:"foreignkey:DocumentTypesId;" json:"document_type" form:"suborders_document_type"`
	DeliveryMethodsId int             `gorm:"column:delivery_methods_id;not null;" json:"delivery_methods_id" form:"suborders_delivery_methods_id"`
	DeliveryMethod    DeliveryMethods `gorm:"foreignkey:DeliveryMethodsId;" json:"delivery_method" form:"suborders_delivery_method"`
	Packages          []Packages      `gorm:"foreignkey:SubordersId;association_foreignkey:Id;" json:"packages"`
}

func (s *Suborders) Expand(data *gorm.DB) error {
	if err := data.Model(s).Related(&s.DocumentType).Error; err != nil {
		return utils.NewError(err, "document type")
	}

	if err := data.Model(s).Related(&s.DeliveryMethod).Error; err != nil {
		return utils.NewError(err, "delivery method")
	}

	if err := data.Model(s).Related(&s.Packages).Error; err != nil {
		if err.Error() == "record not found" {
			return nil
		} else {
			return utils.NewError(err, "packages")
		}
	} else {
		for i := range s.Packages {
			if err := s.Packages[i].Expand(data); err != nil {
				return utils.NewError(err, "packages expand")
			}
		}
	}

	return nil
}

func (s *Suborders) GetSuborderByCode(data *gorm.DB) bool {
	if err := data.Where("code = ?", s.Code).Find(&s).Error; err != nil {
		return false
	} else {
		return true
	}
}

func (s *Suborders) GetVolume() float64 {
	if len(s.Packages) > 0 {
		var volume float64
		for i := range s.Packages {
			volume += s.Packages[i].GetVolume()
		}
		return volume
	} else {
		return 0
	}
}

func (s *Suborders) GetWeight() float64 {
	if len(s.Packages) > 0 {
		var weight float64
		for i := range s.Packages {
			weight += s.Packages[i].Weight
		}
		return weight
	} else {
		return 0
	}
}
