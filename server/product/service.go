package product

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const ProductsPath = "products"

type productService struct {
	repo *ProductRepo
}

func (svc *productService) handleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		products, err := (*svc.repo).List() //getProductList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(products)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		//TODO:
		return
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func SetupRoutes(apiBasePath string, repo ProductRepo) {
	service := productService{repo: &repo}

	productsHandler := http.HandlerFunc(service.handleProducts)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, ProductsPath), productsHandler)
}
