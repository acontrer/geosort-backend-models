package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Routes struct {
	Id                    int                 `gorm:"column:id;not null;" json:"id" form:"routes_id"`
	IdTrl                 int                 `gorm:"column:id_trl;not null;" json:"id_trl" form:"routes_id_trl"`
	DriversId             *int                `gorm:"column:drivers_id;" json:"drivers_id" form:"routes_drivers_id"`
	Driver                Drivers             `gorm:"foreignkey:DriversId;" json:"driver"`
	VehiclesId            *int                `gorm:"column:vehicles_id;" json:"vehicles_id" form:"routes_vehicles_id"`
	Vehicle               Vehicles            `gorm:"foreignkey:VehiclesId;" json:"vehicle" form:"routes_vehicle"`
	DistributionCentersId int                 `gorm:"column:distribution_centers_id;not null;" json:"distribution_centers_id" form:"routes_distribution_centers_id"`
	DistributionCenter    DistributionCenters `gorm:"foreignkey:DistributionCentersId;" json:"distribution_center" form:"routes_distribucion_center"`
	StartTime             time.Time           `gorm:"column:start_time;not null;" json:"start_time" form:"routes_start_time"`
	FinishTime            time.Time           `gorm:"column:finish_time;not null;" json:"finish_time" form:"routes_finish_time"`
	EstimatedFinishTime   time.Time           `gorm:"column:estimated_finish_time;not null;" json:"estimated_finish_time" form:"routes_estimated_finish_time"`
	NextPoint             *int64              `gorm:"column:next_point;" json:"next_point" form:"routes_next_point"`
	DeliveryPoints        []DeliveryPoints    `gorm:"foreignkey:RoutesId;association_foreignkey:Id;" json:"delivery_points" form:"routes_delivery_points"`
	TravelTypesId         int                 `gorm:"column:travel_types_id;not null;" json:"travel_types_id" form:"routes_travel_types_id"`
	TravelType            TravelTypes         `gorm:"foreignkey:TravelTypesId;" json:"travel_types" form:"routes_travel_types"`
	Value                 float64             `gorm:"column:value;not null;" json:"value" form:"routes_value"`
}

func (r *Routes) GetRouteByIdTrl(db *gorm.DB, idTRL int) error {
	if err := db.Where("id_trl = ?", idTRL).Find(&idTRL).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func (r *Routes) GetLastPoint() *DeliveryPoints {
	return &r.DeliveryPoints[len(r.DeliveryPoints)-1]
}

func (r *Routes) TransferRoute(position, idVehicles int) Routes {
	nextPoint := new(int64)
	*nextPoint = r.DeliveryPoints[position].Id

	copyIdVehicles := new(int)
	*copyIdVehicles = idVehicles

	var newRoute = Routes{
		DistributionCentersId: r.DistributionCentersId,
		StartTime:             r.StartTime,
		NextPoint:             nextPoint,
		VehiclesId:            copyIdVehicles,
	}
	allDeliveryPoints := make([]DeliveryPoints, len(r.DeliveryPoints))
	copy(allDeliveryPoints, r.DeliveryPoints)

	newRoute.DeliveryPoints = make([]DeliveryPoints, len(allDeliveryPoints[position:]))
	copy(newRoute.DeliveryPoints, r.DeliveryPoints[position:])

	r.DeliveryPoints = make([]DeliveryPoints, len(allDeliveryPoints[:position]))
	copy(r.DeliveryPoints, allDeliveryPoints[:position])

	return newRoute
}

func (r *Routes) Expand(data *gorm.DB) error {
	if err := data.Model(r).Related(&r.DeliveryPoints).Error; err != nil {
		return err
	} else {
		for i := range r.DeliveryPoints {
			if err := r.DeliveryPoints[i].Expand(data); err != nil {
				return err
			}
		}
	}

	if err := data.Model(r).Related(&r.DistributionCenter).Error; err != nil {
		return err
	} else {
		if err := r.DistributionCenter.Expand(data); err != nil {
			return err
		}
	}

	if err := data.Model(r).Related(&r.Driver).Error; err != nil {
		return err
	} else {

		if err := r.Driver.Expand(data); err != nil {
			return err
		}
	}
	if err := data.Model(r).Related(&r.Vehicle).Error; err != nil {
		return err
	} else {
		if err := r.Vehicle.Expand(data); err != nil {
			return err
		}
	}

	return nil
}
