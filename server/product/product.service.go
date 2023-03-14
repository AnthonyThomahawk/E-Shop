package product

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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

func (svc *productService) handleProduct(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		svc.getProductByID(w, r)
	case http.MethodPost:
		//TODO:
		return
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (svc *productService) getProductByID(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", ProductsPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	productID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	products, err := (*svc.repo).Details(productID) //getProductList()
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

func (svc *productService) getProducts(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()

	page, err := strconv.Atoi(queries.Get("page"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if page == 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(queries.Get("page_size"))
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

	var categoryID *int
	if rawCategoryID := queries.Get("category"); rawCategoryID != "" {
		categoryID = new(int)
		*categoryID, err = strconv.Atoi(rawCategoryID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	products, err := (*svc.repo).List(page, pageSize, categoryID) //getProductList()
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
	productHandler := http.HandlerFunc(service.handleProduct)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, ProductsPath), cors.Middleware(productsHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, ProductsPath), cors.Middleware(productHandler))
}
