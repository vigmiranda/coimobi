package service

import (
	"github.com/vigmiranda/coimobi-service/internal/property/model"
	"github.com/vigmiranda/coimobi-service/pkg/configs"
	"time"
)

func GetAllProperties() []model.Property {
	var properties []model.Property
	configs.DB.Find(&properties).Where("deleted_at = ?", nil)
	return properties
}
func GetProperty(id uint) model.Property {
	var property model.Property
	configs.DB.Find(&property).Where("id = ?", id)
	return property
}

func InsertProperty(property model.Property) model.Property {
	configs.DB.Create(&property)
	return property
}

func UpdateProperty(id uint, updatedData model.Property) model.Property {
	var property model.Property
	configs.DB.Model(&property).Where("id = ?", id).Updates(updatedData)
	return property
}

func SoftDeleteProperty(id uint) error {
	return configs.DB.Model(&model.Property{}).
		Where("id = ?", id).
		Update("deleted_at", time.Now()).Error
}
