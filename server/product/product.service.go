package product

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AnthonyThomahawk/E-Shop/server/cors"
)

const ProductsPath = "products"

type productService struct {
	repo *ProductRepo
}

func (svc *productService) handleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		svc.getProducts(w, r)
	case http.MethodPost:
		//TODO:
		return
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (svc *productService) getProducts(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	page, err := strconv.Atoi(q.Get("page"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if page == 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(q.Get("page_size"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	products, err := (*svc.repo).List(page, pageSize) //getProductList()
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
}

func SetupProductRoutes(apiBasePath string, repo ProductRepo) {
	service := productService{repo: &repo}

	productsHandler := http.HandlerFunc(service.handleProducts)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, ProductsPath), cors.Middleware(productsHandler))
	fmt.Printf("%s/%s\n", apiBasePath, ProductsPath)
}
