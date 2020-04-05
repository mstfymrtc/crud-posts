package app

import (
	"context"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mstfymrtc/go-posts-api/models"
	u "github.com/mstfymrtc/go-posts-api/utils"
	"net/http"
	"os"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api/register", "/api/authenticate"} //don't require authentication
		requestPath := r.URL.Path                                 //current request path

		//if current path is one of the ones that does not require auth, serve request
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")

		if (tokenHeader == "") {
			response = u.Message(false, "Auth token can't be found")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //check if token in format: Bearer {token}
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}
		tokenPart := splitted[1]
		tk := &models.Token{}

		//decode token with token_password
		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (i interface{}, e error) {
			return []byte(os.Getenv("token_password")), nil
		})
		//an error occured while decoding the token
		if err != nil {
			response = u.Message(false, "Broken token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		//token is invalid because of various reasons
		if (!token.Valid) {
			response = u.Message(false, "Invalid token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		//everything is ok, token is valid.
		fmt.Sprintf("User %", tk.UserId)
		ctx := context.WithValue(r.Context(), "user_id", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
