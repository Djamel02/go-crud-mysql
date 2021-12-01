package middlewares

import (
	env "crud/utils"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

type Jsonerror struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}



func IsAuthorized(cb func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	response, _ := json.Marshal(
		&Jsonerror{
			Success: false,
			Message: "Unauthorized",
		},
	)
	return func(w http.ResponseWriter, r *http.Request) {
		jwtToken := strings.Split(r.Header.Get("Authorization"),"Bearer ") ;
		if len(jwtToken) == 1 || jwtToken[1] == "" {
			
			w.Header().Set("Content-Type","application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(response))
			return
		}

		// Decoding token
		token, err := jwt.Parse(jwtToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("jwt signin method error")
			}
			return []byte(env.GetEnvironmentVars("JWT_SECRET")), nil
		})
		if err != nil {
			w.Header().Set("Content-Type","application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		if token.Valid {
			cb(w, r)
		} else {
			w.Header().Set("Content-Type","application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(response))
			return
		}

	}
}