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

const CartPath = "cart"

type cartService struct {
	repo *CartRepo
}

func (svc *cartService) handleCart(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		svc.getCart(w, r)
	case http.MethodPost:
		svc.insertCartItem(w, r)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (svc *cartService) handleCartItem(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		svc.updateCartItem(w, r)
	case http.MethodDelete:
		svc.deleteCartItem(w, r)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (svc *cartService) insertCartItem(w http.ResponseWriter, r *http.Request) {
	var model = struct {
		ProductID uint `json:"ProductId"`
		Quantity  uint
	}{}
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//TODO: Extract UserID
	userID := uint(1)

	err = (*svc.repo).Insert(userID, model.ProductID, model.Quantity)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func (svc *cartService) updateCartItem(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", CartPath))
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

	var model = struct {
		Quantity uint
	}{}
	err = json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//TODO: Extract UserID
	userID := uint(1)

	err = (*svc.repo).Update(userID, uint(productID), model.Quantity)
	if err != nil {
		if err.Error() == "record not found" {
			log.Print(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (svc *cartService) deleteCartItem(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", CartPath))
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

	//TODO: Extract UserID
	userID := uint(1)

	err = (*svc.repo).Delete(userID, uint(productID))
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (svc *cartService) getCart(w http.ResponseWriter, r *http.Request) {
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

	// TODO: Extract UserID form token
	var userID uint = 1

	cartRows, err := (*svc.repo).List(page, pageSize, userID) //getProductList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cart := prepareCart(cartRows)

	j, err := json.Marshal(cart)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(j)
	if err != nil {
		log.Fatal(err)
	}
}

func prepareCart(cartRows []UserProductCart) Cart {
	cart := Cart{
		Items: make([]CartItem, 0, len(cartRows)),
	}
	for _, row := range cartRows {
		cart.Items = append(cart.Items, CartItem{
			Product:  row.Product,
			Quantity: row.Quantity,
		})
	}
	return cart
}

func SetupCartRoutes(apiBasePath string, repo CartRepo) {
	service := cartService{repo: &repo}
	cartHandler := http.HandlerFunc(service.handleCart)
	cartItemHandler := http.HandlerFunc(service.handleCartItem)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, CartPath), cors.Middleware(cartHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, CartPath), cors.Middleware(cartItemHandler))
}
