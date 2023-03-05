package product


import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AnthonyThomahawk/E-Shop/server/cors"
)

const CategoriesPath = "categories"

type categoryService struct {
	repo *CategoryRepo
}

func (svc *categoryService) handleCategories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		svc.getCategories(w, r)
	case http.MethodPost:
		//TODO:
		return
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (svc *categoryService) getCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := (*svc.repo).List() //getProductList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	j, err := json.Marshal(categories)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(j)
	if err != nil {
		log.Fatal(err)
	}
}

func SetupCategoryRoutes(apiBasePath string, repo CategoryRepo) {
	service := categoryService{repo: &repo}

	categoryHandler := http.HandlerFunc(service.handleCategories)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, CategoriesPath), cors.Middleware(categoryHandler))
	fmt.Printf("%s/%s\n", apiBasePath, CategoriesPath)
}
