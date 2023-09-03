package database

import (
	"context"
	"fmt"
	"hbcase/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true&interpolateParams=true", config.GetFromEnv("MYSQL_USER"),
		config.GetFromEnv("MYSQL_PASSWORD"),
		config.GetFromEnv("MYSQL_HOST"),
		config.GetFromEnv("MYSQL_DATABASE"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Println(err)
		return err
	}

	DB = db.WithContext(context.TODO())

	fmt.Println("Connected to database")
	return nil
}
