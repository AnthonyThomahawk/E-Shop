package main

import (
	"log"
	"net/http"

	"github.com/AnthonyThomahawk/E-Shop/server/auth"
	"github.com/AnthonyThomahawk/E-Shop/server/database"
	"github.com/AnthonyThomahawk/E-Shop/server/product"
)

const BasePath = "/api"

func main() {
	db, err := database.Setup()
	if err != nil {
		log.Fatal(err)
	}

	err = auth.SetupAuth()
	if err != nil {
		log.Fatal(err)
	}

	productRepo := database.NewProductRepo(db)
	product.SetupProductRoutes(BasePath, productRepo)

	categoryRepo := database.NewCategoryRepo(db)
	product.SetupCategoryRoutes(BasePath, categoryRepo)

	cartRepo := database.NewCartRepo(db)
	product.SetupCartRoutes(BasePath, cartRepo)

	userRepo := database.NewUserRepo(db)
	auth.SetupLoginRoutes(BasePath, userRepo)
	auth.SetupRegisterRoutes(BasePath, userRepo)

	log.Println("Listening to 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
