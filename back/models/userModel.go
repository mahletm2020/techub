package models
import (
	"gorm.io/gorm"
)

type User struct {
  gorm.Model
  UserName string 
	email    string   `gorm:"uniqe"`
	paswword string 
	role     string
}