package service

import (
	"github.com/gin-gonic/gin"
	"github.com/vigmiranda/coimobi-service/internal/auth/controller"
	"github.com/vigmiranda/coimobi-service/internal/auth/model"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if !Authenticate(req.Email, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	token, err := controller.GenerateJWT(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, model.LoginResponse{Token: token})
}
