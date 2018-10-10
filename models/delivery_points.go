package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"googlemaps.github.io/maps"
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
	Distance              float64             `gorm:"column:distance;" json:"distance" form:"delivery_points_distance"`
	TimeWindows           []TimeWindows       `gorm:"many2many:time_details;" json:"time_windows" form:"delivery_points_time_windows"`
	TravelTypesId         int                 `gorm:"column:travel_types_id;not null;" json:"travel_types_id" form:"delivery_points_travel_types_id"`
	TravelType            TravelTypes         `gorm:"foreignkey:TravelTypesId;" json:"travel_types" form:"delivery_points_travel_types"`
}

func (dp *DeliveryPoints) Expand(data *gorm.DB) error {
	if err := data.Model(dp).Related(&dp.Point).Error; err != nil {
		return err
	} else {
		if err := dp.Point.Expand(data); err != nil {
			return err
		}
	}

	if err := data.Model(dp).Related(&dp.Customer).Error; err != nil {
		return err
	}

	if err := data.Model(dp).Related(&dp.DistributionCenter).Error; err != nil {
		return err
	}

	if err := data.Model(dp).Related(&dp.Suborders).Error; err != nil {
		return err
	} else {
		for i := range dp.Suborders {
			if err := dp.Suborders[i].Expand(data); err != nil {
				return err
			}
		}
	}

	if err := data.Model(dp).Related(&dp.TimeWindows, "TimeWindows").Error; err != nil {
		return nil
	}

	if err := data.Model(dp).Related(&dp.TravelType).Error; err != nil {
		return err
	}

	return nil
}

func (dp DeliveryPoints) GetLatLng() maps.LatLng {
	return maps.LatLng{
		Lat: dp.Point.Latitude,
		Lng: dp.Point.Longitude,
	}
}
