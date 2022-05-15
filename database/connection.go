package database

import (
	"goadmin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	database, err := gorm.Open(mysql.Open("root:root@123@/goadmin"), &gorm.Config{})
	if err != nil {
		panic("db failed to connect")
	}
	DB = database
	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{},models.OrderItem{})
}
