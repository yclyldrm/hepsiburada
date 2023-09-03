package app

import (
	"hbcase/config"
	"hbcase/internal/database"
	"hbcase/internal/domain/repository"
	"hbcase/internal/domain/services"
	"hbcase/pkg"
	"log"
)

func StartApp() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config.LoadConfig(".env")
	if err := database.InitDB(); err != nil {
		log.Fatal(err)
	}
	database.Migrate()

	ps := services.NewProductService(repository.NewProductRepository(database.DB))
	os := services.NewOrderService(repository.NewOrderRepository(database.DB))
	cs := services.NewCampaignService(repository.NewCampaignRepository(database.DB))

	cmdService = pkg.NewCommandService(ps, os, cs)

	router := NewRouter()

	router.Run(":" + config.GetFromEnv("PORT"))
}
