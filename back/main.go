package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var secretKey string

// Initialize environment variables and secret key
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey = os.Getenv("Secretkey")
	if secretKey == "" {
		log.Fatal("Secretkey environment variable not set")
	}
}

// Function to hash password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Function to create JWT token
func createToken(username, email, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"email":    email,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}

// Registration handler
func register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Role     string `json:"role" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password encryption failed"})
		return
	}

	token, err := createToken(input.Username, input.Email, input.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token creation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "hashed_password": hashedPassword})
}

// Main function
func main() {
	r := gin.Default()
	r.POST("/register", register)

	port := os.Getenv("Port")
	if port == "" {
		port = "3088"
	}
	r.Run(":" + port)
}
