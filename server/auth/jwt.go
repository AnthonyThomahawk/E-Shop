package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secretyoushouldhide")

type Claims struct {
	UserID uint //`json:"UserId"`
	RoleID uint //`json:"RoleId"`
	jwt.RegisteredClaims
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
	//SigningMethodEdDSA didn't work
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(endpointHandler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})
}

// func extractClaims(_ http.ResponseWriter, request *http.Request) (string, error) {
// 	if request.Header["Token"] != nil {
// 		tokenString := request.Header["Token"][0]
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//
// 			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
// 				return nil, fmt.Errorf("there's an error with the signing method")
// 			}
// 			return secretKey, nil
// 		})
//
// 		if err != nil {
// 			return "Error Parsing Token: ", err
// 		}
// 	}
// }

func ExtractClaims(request *http.Request) (*Claims, error) {
	if tokenString := request.Header.Get("Authorization"); tokenString != "" {
		tokenString= strings.Split(tokenString, " ")[1]
		token, err := jwt.ParseWithClaims(tokenString, &Claims{},func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there's an error with the signing method")
			}
			return secretKey, nil
		})
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
