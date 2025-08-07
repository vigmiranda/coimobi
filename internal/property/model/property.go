package model

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	Title        string  `json:"title"`          // Ex: "3-Bedroom Apartment - Moema/SP"
	PropertyType string  `json:"property_type"`  // Apartment, House, Land, Commercial, etc.
	Purpose      int64   `json:"purpose"`        // Sale, Rent, Income, Resale
	Description  string  `json:"description"`    // Full description, highlights, location details
	UsableAreaM2 float64 `json:"usable_area_m2"` // Usable area in square meters
	TotalAreaM2  float64 `json:"total_area_m2"`  // Total area in square meters
	Bedrooms     int     `json:"bedrooms"`       // Including suites
	Bathrooms    int     `json:"bathrooms"`
	GarageSpaces int     `json:"garage_spaces"`
	FullAddress  string  `json:"full_address"` // Street, number, neighborhood, city, state, ZIP
	Latitude     float64 `json:"latitude"`     // For map integration
	Longitude    float64 `json:"longitude"`
}
