package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TimeWindows struct {
	Id           int            `gorm:"column:id;not null;" json:"id" form:"time_windows_id"`
	TimeInit     time.Time      `gorm:"column:time_init;not null;" json:"time_init" form:"time_windows_time_init" time_format:"2006-01-02T15:04:05Z"`
	TimeFinish   time.Time      `gorm:"column:time_finish;not null;" json:"time_finish" form:"time_windows_time_finish" time_format:"2006-01-02T15:04:05Z"`
	Restrictions []Restrictions `gorm:"foreignkey:time_windows_id;" json:"restrictions" form:"time_windows_restrictions"`
}

func (tw *TimeWindows) GetID(db *gorm.DB, timeInit, timeFinal time.Time) bool {
	//db, err := db.Database()
	//defer db.Close()
	//
	//if err != nil {
	//	return false
	//} else {
	if err := db.Where("time_init = ? AND time_finish = ?", timeInit, timeFinal).Find(&tw).Error; err != nil {
		return false
	} else {
		return true
	}
	//}
}

func (tw *TimeWindows) Expand(data *gorm.DB) error {
	if err := data.Model(tw).Related(&tw.Restrictions).Error; err != nil {
		return err
	} else {
		for i, _ := range tw.Restrictions {
			if err := tw.Restrictions[i].Expand(data); err != nil {
				return err
			}
		}
	}
	return nil
}

func (tw *TimeWindows) IsInside(t time.Time) bool {
	auxinit := time.Date(2000, time.January, 1, tw.TimeInit.Hour(), tw.TimeInit.Minute(), 0, 0, time.UTC)
	auxfinish := time.Date(2000, time.January, 1, tw.TimeFinish.Hour(), tw.TimeFinish.Minute(), 0, 0, time.UTC)
	auxt := time.Date(2000, time.January, 1, t.Hour(), t.Minute(), 0, 0, time.UTC)
	return auxt.After(auxinit) && auxt.Before(auxfinish)
}
