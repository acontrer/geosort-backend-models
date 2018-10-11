package models

import (
	"strings"

	"github.com/dwladdimiroc/geosort-backend-models/utils"
	"github.com/jinzhu/gorm"
)

type Streets struct {
	Id         int      `gorm:"column:id;not null;" json:"id" form:"streets_id"`
	CommunesId int      `gorm:"column:communes_id;not null;" json:"communes_id" form:"streets_communes_id"`
	Commune    Communes `gorm:"foreignkey:CommunesId" json:"commune" form:"streets_commune"`
	Name       string   `gorm:"column:name;not null;" json:"name" form:"streets_name"`
}

func (s *Streets) Expand(data *gorm.DB) error {
	if err := data.Model(s).Related(&s.Commune).Error; err != nil {
		return err
	} else {
		if err := s.Commune.Expand(data); err != nil {
			return err
		}
	}

	return nil
}

func (s *Streets) GetStreetByName(db *gorm.DB, name string) bool {
	name = strings.ToLower(utils.ValidString(name))
	if err := db.Where("lower(name) = ?", name).Find(&s).Error; err != nil {
		return false
	} else {
		return true
	}
}

func (s *Streets) GetCommune(db *gorm.DB, commune, region string) bool {
	region = strings.ToLower(utils.ValidString(region))
	if err := db.Where("lower(name) = ?", region).Find(&s.Commune.Region).Error; err != nil {
		return false
	} else {
		commune = strings.ToLower(utils.ValidString(commune))
		if err := db.Where("lower(name) =  ? AND regions_id = ?", commune, s.Commune.Region.Id).Find(&s.Commune).Error; err != nil {
			return false
		} else {
			s.CommunesId = s.Commune.Id
			//TODO Cbange zones id
			s.Commune.ZonesId = 1
			s.Commune = Communes{}
			return true
		}
	}
}
