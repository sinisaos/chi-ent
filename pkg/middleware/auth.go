package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/sinisaos/chi-ent/pkg/database"

	"github.com/golang-jwt/jwt/v5"
)

var privateKey = []byte(database.Config("SECRET_KEY"))

// AuthMiddleware to authenticate users
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return privateKey, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}

		userID := token.Claims.(jwt.MapClaims)["user_id"].(float64)
		headerUserID := strconv.FormatFloat(userID, 'f', -1, 64)
		r.Header.Set("user_id", headerUserID)

		next.ServeHTTP(w, r)
	})
}
