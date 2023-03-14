package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AnthonyThomahawk/E-Shop/server/cors"
	"github.com/AnthonyThomahawk/E-Shop/server/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const RegisterPath = "auth/register"

type registerService struct {
	repo *user.UserRepo
}

func (svc *registerService) handleRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		svc.register(w, r)
		return
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (svc *registerService) register(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usr, err := (*svc.repo).Load(creds.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if usr != nil {
		//TODO: add message "there is already user with that email"
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		// TODO: add message "could not hash password; %v", err
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUser := user.User{
		Email:        creds.Email,
		PasswordHash: string(hash[:]),
		RoleID:       2,
	}

	err = (*svc.repo).Save(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := generateJWT(newUser.Email, newUser.RoleID)
	if err != nil {
		return
	}

	j, err := json.Marshal(Session{
		Email: creds.Email,
		Token: token,
	})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = w.Write(j)
	if err != nil {
		log.Fatal(err)
	}
}

func SetupRegisterRoutes(apiBasePath string, repo user.UserRepo) {
	service := registerService{repo: &repo}
	registerHandler := http.HandlerFunc(service.handleRegister)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, RegisterPath), cors.Middleware(registerHandler))
}
