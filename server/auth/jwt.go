package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey []byte

type Claims struct {
	UserID uint
	RoleID uint
	jwt.RegisteredClaims
}

func SetupAuth() error {
	if sk, found := os.LookupEnv("AUTH_SECRET_KEY"); found {
		secretKey = []byte(sk)
		return nil
	}

	return errors.New("AUTH_SECRET_KEY not provided")
}

func GenerateJWT(userID, roleID uint) (string, error) {
	claims := Claims{
		UserID: userID,
		RoleID: roleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(20 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(endpointHandler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ignoring OPTION preflight request
		if r.Method == http.MethodOptions {
			return
		}

		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString = strings.Split(tokenString, " ")[1]

		token, err := jwt.ParseWithClaims(
			tokenString,
			&Claims{},
			func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					w.WriteHeader(http.StatusUnauthorized)
					_, err := w.Write([]byte("You're Unauthorized!"))
					if err != nil {
						return nil, err
					}
				}
				return secretKey, nil
			},
		)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if _, ok := token.Claims.(*Claims); !ok || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		endpointHandler(w, r)
	}
}

func ExtractClaims(request *http.Request) (*Claims, error) {
	if tokenString := request.Header.Get("Authorization"); tokenString != "" {
		tokenString = strings.Split(tokenString, " ")[1]
		token, err := jwt.ParseWithClaims(
			tokenString,
			&Claims{},
			func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there's an error with the signing method")
				}
				return secretKey, nil
			},
		)
		if err != nil {
			return nil, err
		}
		if token.Valid {
			if claims, ok := token.Claims.(*Claims); ok {
				return claims, nil
			}
		}
	}

	return nil, nil
}
