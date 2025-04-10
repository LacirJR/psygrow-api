package main

import (
	"github.com/LacirJR/psygrow-api/src/internal/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	app := gin.Default()
	router.RegisterRoutes(app)

	log.Printf("Iniciando servidor na porta 8080...")

	err := app.Run(":8080")
	if err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}

}
