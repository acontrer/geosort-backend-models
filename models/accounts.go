package models

import (
	"github.com/jinzhu/gorm"
)

type Accounts struct {
	Id             int          `gorm:"column:id;not null;" json:"id" form:"accounts_id"`
	Password       string       `gorm:"column:password;not null;" json:"password" form:"accounts_password"`
	Email          string       `gorm:"column:email;not null;" json:"email" form:"accounts_email"`
	AccountTypesId int          `gorm:"column:account_types_id;not null;" json:"account_types_id" form:"accounts_account_types_id"`
	AccountType    AccountTypes `gorm:"foreignkey:AccountTypesId;" json:"account_type" form:"account_account_type"`
}

func (a *Accounts) Expand(data *gorm.DB) error {
	if err := data.Model(a).Related(&a.AccountType).Error; err != nil {
		return err
	}

	return nil
}

func (a *Accounts) GetAccountByEmail(data *gorm.DB, email string) error {
	if err := data.Where("email = ?", email).First(a).Error; err != nil {
		return err
	} else {
		return nil
	}
}
