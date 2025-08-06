package service

import (
	"github.com/vigmiranda/coimobi-service/api/configs"
	"github.com/vigmiranda/coimobi-service/api/models"
	"time"
)

func GetAllProperties() []models.Property {
	var properties []models.Property
	configs.DB.Find(&properties).Where("deleted_at = ?", nil)
	return properties
}

func InsertProperty(property models.Property) models.Property {
	configs.DB.Create(&property)
	return property
}

func UpdateProperty(id uint, updatedData models.Property) models.Property {
	var property models.Property
	configs.DB.Model(&property).Where("id = ?", id).Updates(updatedData)
	return property
}

func SoftDeleteProperty(id uint) error {
	return configs.DB.Model(&models.Property{}).
		Where("id = ?", id).
		Update("deleted_at", time.Now()).Error
}
