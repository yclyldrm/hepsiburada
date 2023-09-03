package database

import (
	"hbcase/internal/domain/entity"
	"log"
)

func Migrate() {
	if err := DB.AutoMigrate(
		entity.Product{},
		entity.Campaign{},
		entity.Order{},
	); err != nil {
		log.Println(err.Error())
	}
}
