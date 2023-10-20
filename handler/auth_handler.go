package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sinisaos/chi-ent/database"
	"github.com/sinisaos/chi-ent/model"
	"github.com/sinisaos/chi-ent/service"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: service,
	}
}

// Login Handler
func (h AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	payload := new(model.LoginUserInput)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	// Check email
	newUser, err := h.AuthService.Login(payload)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(payload.Password)); err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
	}

	// Create a token for the user with the correct email and password
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = newUser.Username
	claims["user_id"] = newUser.ID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	jwtToken, err := token.SignedString([]byte(database.Config("SECRET_KEY")))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondwithJSON(w, http.StatusOK,
		map[string]interface{}{
			"message": "Successfully logged in",
			"token":   jwtToken,
		},
	)
}
