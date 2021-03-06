package models

import (
	"sort"
	"time"

	"github.com/dwladdimiroc/geosort-backend-models/utils"

	"github.com/jinzhu/gorm"
)

type DeliveryPoints struct {
	Id                    int64               `gorm:"column:id;not null;" json:"id" form:"delivery_points_id"`
	RoutesId              *int                `gorm:"column:routes_id;" json:"routes_id" form:"delivery_points_routes_id"`
	PointsId              int                 `gorm:"column:points_id;not null;" json:"points_id" form:"delivery_points_points_id"`
	Point                 Points              `gorm:"foreignkey:PointsId;" json:"point" form:"delivery_points_point"`
	AdditionalInfoPoint   string              `gorm:"column:additional_info_point;" json:"additional_info_point" form:"delivery_points_additional_info_point"`
	RawAddress            string              `gorm:"column:raw_address;not null;" json:"raw_address" form:"delivery_points_raw_address"`
	CustomersId           int                 `gorm:"column:customers_id;not null;" json:"customers_id" form:"delivery_points_customers_id"`
	Customer              Customers           `gorm:"foreignkey:CustomersId;" json:"customer" form:"delivery_point_customer"`
	RoutePosition         *int                `gorm:"column:route_position;not null;" json:"route_position" form:"delivery_points_route_position"`
	DistributionCentersId int                 `gorm:"column:distribution_centers_id;not null;" json:"distribution_centers_id" form:"delivery_points_distribution_centers_id"`
	DistributionCenter    DistributionCenters `gorm:"foreignkey:DistributionCentersId;" json:"distribution_center" form:"delivery_points_distribution_center"`
	ArrivalAt             time.Time           `gorm:"column:arrival_at;" json:"arrival_at" form:"delivery_points_arrival_at"`
	EstimatedAt           time.Time           `gorm:"column:estimated_at;" json:"estimated_at" form:"delivery_points_estimated_at"`
	CreatedAt             time.Time           `gorm:"column:created_at;not null;" json:"created_at" form:"delivery_points_created_at"`
	Latitude              float64             `gorm:"column:latitude;" json:"latitude" form:"delivery_points_latitude"`
	Longitude             float64             `gorm:"column:longitude;" json:"longitude" form:"delivery_points_longitude"`
	Suborders             []Suborders         `gorm:"foreignkey:DeliveryPointsId;association_foreignkey:Id;" json:"suborders" form:"delivery_points_suborders"`
	TimeWindows           []TimeWindows       `gorm:"many2many:time_details;" json:"time_windows" form:"delivery_points_time_windows"`
	TravelTypesId         int                 `gorm:"column:travel_types_id;not null;" json:"travel_types_id" form:"delivery_points_travel_types_id"`
	TravelType            TravelTypes         `gorm:"foreignkey:TravelTypesId;" json:"travel_types" form:"delivery_points_travel_types"`
	RatioLatitude         float64             `gorm:"column:ratio_latitude;not null;" json:"ratio_latitude" form:"delivery_points_ratio_latitude"`
	RatioLongitude        float64             `gorm:"column:ratio_longitude;not null;" json:"ratio_longitude" form:"delivery_points_ratio_longitude"`
	RatioArrivalAt        time.Time           `gorm:"column:ratio_arrival_at;not null;" json:"ratio_arrival_at" form:"delivery_points_ratio_arrival_at"`
	RequestedAt           time.Time           `gorm:"column:requested_at;not null;" json:"ratio_arrival_at" form:"delivery_points_ratio_arrival_at"`
}

func (dp *DeliveryPoints) Expand(data *gorm.DB) error {
	if err := data.Model(dp).Related(&dp.Point).Error; err != nil {
		return utils.NewError(err, "point")
	} else {
		if err := dp.Point.Expand(data); err != nil {
			return utils.NewError(err, "point expand")
		}
	}

	if err := data.Model(dp).Related(&dp.Customer).Error; err != nil {
		return utils.NewError(err, "customer")
	}

	if err := data.Model(dp).Related(&dp.DistributionCenter).Error; err != nil {
		return utils.NewError(err, "distribution center")
	}

	if err := data.Model(dp).Related(&dp.TravelType).Error; err != nil {
		return utils.NewError(err, "travel type")
	}

	if err := data.Model(dp).Related(&dp.Suborders).Error; err != nil {
		return utils.NewError(err, "suborder")
	} else {
		for i := range dp.Suborders {
			if err := dp.Suborders[i].Expand(data); err != nil {
				return utils.NewError(err, "suborder expand")
			}
		}
	}

	if err := data.Model(dp).Related(&dp.TimeWindows, "TimeWindows").Error; err != nil {
		return utils.NewError(err, "time windows")
	}

	return nil
}

func (dp *DeliveryPoints) SearchTimeWindow(tw TimeWindows) int {
	idx := sort.Search(len(dp.TimeWindows), func(i int) bool {
		return dp.TimeWindows[i].TimeInit == tw.TimeInit && dp.TimeWindows[i].TimeFinish == tw.TimeFinish
	})
	if idx < len(dp.TimeWindows) && dp.TimeWindows[idx].TimeInit == tw.TimeInit && dp.TimeWindows[idx].TimeFinish == tw.TimeFinish {
		return idx
	}
	return -1
}

func (dp *DeliveryPoints) AddTimeWindowWithoutZone(data *gorm.DB, tw TimeWindows) error {
	if dp.SearchTimeWindow(tw) == -1 {
		if err := data.Model(&dp).Association("TimeWindows").Append(&TimeWindows{Id: tw.Id}).Error; err != nil {
			return err
		}
	}
	return nil
}

func (dp *DeliveryPoints) AddTimeWindow(data *gorm.DB) error {
	if r, err := dp.Point.GetRestrictions(data); err != nil {
		return err
	} else {
		weekday := dp.RequestedAt.Weekday()
		for i, _ := range r {
			if (int(weekday) + 1) == r[i].DaysId {
				if err := data.Model(&dp).Association("TimeWindows").Append(&TimeWindows{Id: r[i].TimeWindowsId}).Error; err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (dp *DeliveryPoints) RemoveTimeWindows(data *gorm.DB) error {
	if len(dp.TimeWindows) > 0 {
		for i, _ := range dp.TimeWindows {
			if err := dp.TimeWindows[i].Expand(data); err != nil {
				return err
			}
			if len(dp.TimeWindows[i].Restrictions) > 0 {
				weekday := dp.RequestedAt.Weekday()
				for j, _ := range dp.TimeWindows[i].Restrictions {
					if (int(weekday) + 1) != dp.TimeWindows[i].Restrictions[j].DaysId {
						if err := data.Model(&dp).Association("TimeWindows").Delete(&TimeWindows{Id: dp.TimeWindows[i].Id}).Error; err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}

func (dp *DeliveryPoints) SetRequestAt(data *gorm.DB, date time.Time) error {
	if err := dp.RemoveTimeWindows(data); err != nil {
		return err
	}
	dp.RequestedAt = date
	if err := dp.AddTimeWindow(data); err != nil {
		return err
	}
	return nil
}
