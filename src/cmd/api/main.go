package main

import (
	"fmt"
	"github.com/LacirJR/psygrow-api/src/internal/config"
	"github.com/LacirJR/psygrow-api/src/internal/infra/migration"
	"github.com/LacirJR/psygrow-api/src/internal/router"
	"github.com/LacirJR/psygrow-api/src/internal/seed"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	//Carregar arquivo .env
	config.LoadEnv()

	//Iniciar banco de dados
	config.InitDatabase()

	//Aplicar migrações
	migration.Migrate()

	//Criar usuario padrao
	seed.CreateDefaultAdminUser()

	//Registrar rotas
	app := gin.Default()
	router.RegisterRoutes(app)

	var port = config.GetEnvironment(config.AppPort)
	//Iniciar servidor
	log.Printf("Iniciando servidor na porta %s...", port)
	err := app.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}

}
