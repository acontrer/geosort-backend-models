package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TimeWindows struct {
	Id         int       `gorm:"column:id;not null;" json:"id" form:"time_windows_id"`
	TimeInit   time.Time `gorm:"column:time_init;not null;" json:"time_init" form:"time_windows_time_init" time_format:"2006-01-02T15:04:05Z"`
	TimeFinish time.Time `gorm:"column:time_finish;not null;" json:"time_finish" form:"time_windows_time_finish" time_format:"2006-01-02T15:04:05Z"`
}

func (tw *TimeWindows) GetID(db *gorm.DB, timeInit, timeFinal time.Time) bool {
	//db, err := db.Database()
	//defer db.Close()
	//
	//if err != nil {
	//	return false
	//} else {
	if err := db.Where("time_init = ? AND time_finish = ?", timeInit, timeFinal).Find(tw).Error; err != nil {
		return false
	} else {
		return true
	}
	//}
}
