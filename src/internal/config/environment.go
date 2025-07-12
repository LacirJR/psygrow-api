package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func GetEnvironment(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Fatalf("Variável de ambiente obrigatória '%s' não está definida", key)
	return ""
}

func LoadEnv() {
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("Aviso: Erro ao obter o diretório atual: %v", err)
	}

	for {
		envPath := filepath.Join(dir, ".env")

		if _, err := os.Stat(envPath); err == nil {
			err = godotenv.Load(envPath)
			if err != nil {
				log.Printf("Aviso: Erro ao carregar .env: %v", err)
			}
			log.Printf(".env carregado de: %s", envPath)
			return
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break // chegou na raiz do sistema
		}
		dir = parent
	}

	log.Printf("Arquivo .env não encontrado em nenhum diretório pai")
}
