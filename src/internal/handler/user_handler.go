package handler

import (
	"github.com/LacirJR/psygrow-api/src/internal/config"
	dto "github.com/LacirJR/psygrow-api/src/internal/core/dto/user"
	"github.com/LacirJR/psygrow-api/src/internal/core/model"
	"github.com/LacirJR/psygrow-api/src/internal/core/security"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// jwtSecretKey will be initialized when needed
var jwtSecretKey []byte

// getJWTSecretKey returns the JWT secret key, initializing it if necessary
func getJWTSecretKey() []byte {
	if len(jwtSecretKey) == 0 {
		jwtSecretKey = []byte(config.GetEnvironment(config.JwtSecretKey))
	}
	return jwtSecretKey
}

func RegisterUser(c *gin.Context) {

	var req dto.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "details": err.Error()})
		return
	}

	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar hash da senha"})
		return
	}

	user := model.User{
		ID:           uuid.New(),
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		Role:         "professional",
		Phone:        req.Phone,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.NewUserResponse(user))
}

func GetUserByEmail(c *gin.Context) {
	email := c.Param("email")

	var user model.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, dto.NewUserResponse(user))

}

func Login(c *gin.Context) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	var user model.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos"})
		return
	}

	if !security.CheckPasswordHash(req.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos"})
		return
	}

	// Check if user is active
	if !user.IsActive {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Conta inativa"})
		return
	}

	// Update last login time
	now := time.Now()
	user.LastLoginAt = &now
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar último login"})
		return
	}

	token, err := generateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func VerifyAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func generateJWT(user model.User) (string, error) {
	// Define as claims do token conforme especificação
	claims := jwt.MapClaims{
		"sub":  user.ID.String(),                      // Subject (ID do usuário)
		"role": user.Role,                             // Papel do usuário
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token válido por 24 horas
	}

	// Cria o token assinado
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJWTSecretKey())
}
