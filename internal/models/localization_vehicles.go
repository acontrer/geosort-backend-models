package models

type LocalizationVehicles struct {
	Id                    int `gorm:"column:id;not null;" json:"id" form:"localization_vehicles_id"`
	DistributionCentersId int `gorm:"column:distribution_centers_id;not null;" json:"distribution_centers_id" form:"localization_vehicles_distribution_centers_id"`
	VehiclesId            int `gorm:"column:vehicles_id;not null;" json:"vehicles_id" form:"localization_vehicles_vehicles_id"`
}
