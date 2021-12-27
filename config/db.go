package config

import (
	"API-REVIEWHP/models"
	"API-REVIEWHP/utils"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "production" {
		username := os.Getenv("DATABASE_USERNAME")
		password := os.Getenv("DATABASE_PASSWORD")
		host := os.Getenv("DATABASE_HOST")
		port := os.Getenv("DATABASE_PORT")
		database := os.Getenv("DATABASE_NAME")
		// production
		dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=require"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		db.AutoMigrate(&models.Brand{}, &models.Devices{}, &models.Specification{}, &models.Reviews{}, &models.Opinions{}, &models.Comments{}, &models.User{})

		return db
	} else {
		// development
		username := "root"
		password := "password"
		host := "tcp(127.0.0.1:3306)"
		database := "db_tugasbesar"

		dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err.Error())
		}

		db.AutoMigrate(&models.Brand{}, &models.Devices{}, &models.Specification{}, &models.Reviews{}, &models.Opinions{}, &models.Comments{}, &models.User{})

		return db
	}

}
