package seed

import (
	"github.com/LacirJR/psygrow-api/src/internal/config"
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/LacirJR/psygrow-api/src/internal/core/security"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
)

func CreateDefaultAdminUser() {
	defaultEmail := config.GetEnvironment("DEFAULT_ADMIN_EMAIL")
	defaultPassword := config.GetEnvironment("DEFAULT_ADMIN_PASSWORD")
	defaultName := config.GetEnvironment("DEFAULT_ADMIN_NAME")
	defaultRole := config.GetEnvironment("DEFAULT_ADMIN_ROLE")

	if defaultEmail == "" || defaultPassword == "" || defaultName == "" {
		log.Println("Aviso: Variáveis de ambiente para o usuário padrão (DEFAULT_ADMIN_EMAIL, DEFAULT_ADMIN_PASSWORD, DEFAULT_ADMIN_NAME) não estão totalmente definidas. Nenhum usuário padrão será criado.")
		return
	}

	// Define um role padrão se não for especificado
	if defaultRole == "" {
		defaultRole = "professional" // Ou "admin", conforme sua lógica de negócio
	}

	// Verifica se o usuário já existe no banco de dados
	var existingUser model.User
	result := config.DB.Where("email = ?", defaultEmail).First(&existingUser)

	if result.Error == nil {
		log.Printf("Usuário padrão com email '%s' já existe. Pulando criação.", defaultEmail)
		return
	}

	if result.Error != gorm.ErrRecordNotFound {
		log.Printf("Erro ao verificar existência do usuário padrão: %v", result.Error)
		return
	}

	// Se o usuário não foi encontrado (ErrRecordNotFound), proceed to create
	hashedPassword, err := security.HashPassword(defaultPassword)
	if err != nil {
		log.Fatalf("Erro fatal ao gerar hash da senha para o usuário padrão: %v", err)
		return // log.Fatalf já encerra o programa
	}

	user := model.User{
		ID:           uuid.New(),
		Name:         defaultName,
		Email:        defaultEmail,
		PasswordHash: hashedPassword,
		Role:         defaultRole,
		Phone:        nil,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		log.Fatalf("Erro fatal ao criar usuário padrão: %v", err)
		return
	}

	log.Printf("Usuário padrão '%s' (%s) criado com sucesso!", user.Name, user.Email)
}
