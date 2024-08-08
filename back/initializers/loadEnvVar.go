package initializers

import (	
	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "github.com/dgrijalva/jwt-go"
	// "golang.org/x/crypto/bcrypt"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

func LoadEnvVar() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
}