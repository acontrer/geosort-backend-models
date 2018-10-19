package models

import (
	"github.com/jinzhu/gorm"
)

type Suborders struct {
	Id                int64      `gorm:"column:id;not null;" json:"id" form:"suborders_id"`
	DeliveryPointsId  int64      `gorm:"column:delivery_points_id;not null;" json:"delivery_points_id" form:"suborders_delivery_points_id"`
	Code              int64      `gorm:"column:code;not null;" json:"code" form:"suborders_code"`
	DocumentTypesId   int        `gorm:"column:document_types_id;not null;" json:"document_types_id" form:"suborders_document_types_id"`
	DeliveryMethodsId int        `gorm:"column:delivery_methods_id;not null;" json:"delivery_methods_id" form:"suborders_delivery_methods_id"`
	Packages          []Packages `gorm:"foreignkey:SubordersId;association_foreignkey:Id;" json:"packages"`
}

func (s *Suborders) Expand(data *gorm.DB) error {
	if err := data.Model(s).Related(&s.Packages).Error; err != nil {
		if err.Error() == "record not found" {
			return nil
		} else {
			return err
		}
	} else {
		for i := range s.Packages {
			if err := s.Packages[i].Expand(data); err != nil {
				return err
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

func(s *Suborders) CheckSku(data *gorm.DB, sku string) bool {
	for i := range s.Packages {
		if s.Packages[i].Sku == sku {
			return true
		}
	}
	return false
}