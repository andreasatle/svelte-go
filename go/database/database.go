package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host string
	Port int
	User string
	Password string
	Name string
}

type User struct {
	gorm.Model
	Name string
	Email string
}

var DB *gorm.DB

// ConnectDB connects to the database
func ConnectDB(config Config) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Name)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	DB.AutoMigrate(&User{})
}