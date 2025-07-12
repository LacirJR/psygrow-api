package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		GetEnvironment(DbHost),
		GetEnvironment(DbUser),
		GetEnvironment(DbPass),
		GetEnvironment(DbName),
		GetEnvironment(DbPort),
		GetEnvironment(SslMode),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados: %v", err)
	}

	DB = db
	log.Println("Banco de dados conectado com sucesso!")

}
