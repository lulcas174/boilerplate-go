package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lulcas174/app-register/database"
)

func Start() {
	database.GetDb()
	// TODO: Adicionar os models
	router := gin.Default()
	// TODO: Criar as rotas

	router.Run(":8000")
}
