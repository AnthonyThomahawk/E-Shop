package auth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secretyoushouldhide")

type Claims struct {
	Email string
	RoleID uint `json:"RoleId"`
	jwt.RegisteredClaims
}

func generateJWT(email string, roleID uint) (string, error) {
	claims:= Claims{
		Email:            email,
		RoleID:           roleID,
		RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(20*time.Minute)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyJWT(endpointHandler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write([]byte("You're Unauthorized!"))
				if err != nil {
					return nil, err
				}
			}
			return "", nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, err2 := w.Write([]byte("You're Unauthorized due to error parsing the JWT"))
			if err2 != nil {
				return
			}
		}

		if token.Valid {
			endpointHandler(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte("You're Unauthorized due to invalid token"))
			if err != nil {
				return
			}
		}
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
