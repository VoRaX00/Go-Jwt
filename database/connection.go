package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"site-cv-back/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "user=nikita password=1324 dbname=auth port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to database")
	}
	DB = connection

	err = connection.AutoMigrate(&models.User{})
}
