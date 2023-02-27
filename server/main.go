package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AnthonyThomahawk/E-Shop/server/config"
	"github.com/AnthonyThomahawk/E-Shop/server/database"
	"github.com/AnthonyThomahawk/E-Shop/server/product"
)

const BasePath = "/api"

func main() {
	cfg, err := config.Import()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Setup(cfg.DB.DataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	productRepo := database.NewProductRepo(db)

	product.SetupRoutes(BasePath, productRepo)
	fmt.Println("Listening to 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
