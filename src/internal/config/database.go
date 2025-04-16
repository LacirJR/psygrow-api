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
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		GetEnvironment(DbHost),
		GetEnvironment(DbUser),
		GetEnvironment(DbPass),
		GetEnvironment(DbName),
		GetEnvironment(DbPort),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados: %v", err)
	}

	DB = db
	log.Println("Banco de dados conectado com sucesso!")

}
