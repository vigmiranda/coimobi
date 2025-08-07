package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vigmiranda/coimobi-service/internal/property/model"
	"github.com/vigmiranda/coimobi-service/internal/property/service"
	"github.com/vigmiranda/coimobi-service/pkg/configs"
	"net/http"
	"strconv"
)

func GetAllProperties(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetAllProperties())
}

func GetProperty(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Buscar o imóvel no banco
	var property model.Property
	if err := configs.DB.First(&property, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Imóvel não encontrado"})
		return
	}
	c.JSON(http.StatusOK, service.GetProperty(uint(id)))
}

func InsertProperty(c *gin.Context) {
	var property model.Property
	if err := c.ShouldBindJSON(&property); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, service.InsertProperty(property))
}

func UpdateProperty(c *gin.Context) {
	// Pegar o ID da URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Buscar o imóvel no banco
	var property model.Property
	if err := configs.DB.First(&property, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Imóvel não encontrado"})
		return
	}

	// Bind JSON recebido para atualizar
	var updatedData model.Property
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Imóvel atualizado com sucesso",
		"property": service.UpdateProperty(property.ID, updatedData),
	})
}

func DeleteProperty(c *gin.Context) {
	// Extrai o ID da URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Busca o imóvel no banco
	var property model.Property
	if err := configs.DB.First(&property, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Imóvel não encontrado"})
		return
	}

	if err := service.SoftDeleteProperty(property.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Imóvel deletado com sucesso"})
}
