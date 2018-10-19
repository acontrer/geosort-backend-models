package models

import "github.com/jinzhu/gorm"

type Enterprises struct {
	Id   int    `gorm:"column:id;not null;" json:"id" form:"enterprises_id"`
	Rut  string `gorm:"column:rut;not null;" json:"rut" form:"enterprises_rut"`
	Name string `gorm:"column:name;not null;" json:"name" form:"enterprises_name"`
}

func GetEnterpriseByRUT(data *gorm.DB, rutEnterprise string) (int, error) {
	var enterprise = Enterprises{
		Rut: rutEnterprise,
	}

	if err := data.Where(enterprise).Find(&enterprise).Error; err != nil {
		return enterprise.Id, err
	} else {
		return enterprise.Id, nil
	}
}
