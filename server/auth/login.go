package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AnthonyThomahawk/E-Shop/server/cors"
	"github.com/AnthonyThomahawk/E-Shop/server/user"
	"golang.org/x/crypto/bcrypt"
)

const LoginPath = "auth/login"

type loginService struct {
	repo *user.UserRepo
}

func (svc *loginService) handleLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		svc.login(w, r)
		return
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (svc *loginService) login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := (*svc.repo).Load(creds.Email)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password)); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := GenerateJWT(user.ID, user.RoleID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
}

func SetupLoginRoutes(apiBasePath string, repo user.UserRepo) {
	service := loginService{repo: &repo}
	loginHandler := http.HandlerFunc(service.handleLogin)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, LoginPath), cors.Middleware(loginHandler))
}
