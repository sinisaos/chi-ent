package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sinisaos/chi-ent/model"
	"github.com/sinisaos/chi-ent/service"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

// All users
func (h UserHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	itemsPerPage, _ := strconv.Atoi(r.URL.Query().Get("itemsPerPage"))
	users, err := h.UserService.GetAllUsers(page, itemsPerPage)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondwithJSON(w, http.StatusOK, map[string]interface{}{"data": users, "page": page})
}

// Single user
func (h UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	user, err := h.UserService.GetUser(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
	}

	respondwithJSON(w, http.StatusOK, user)
}

// New user
func (h UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	payload := new(model.NewUserInput)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	user, err := h.UserService.CreateUser(payload)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondwithJSON(w, http.StatusCreated, user)
}

// Update user
func (h UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	payload := new(model.UpdateUserInput)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	user, err := h.UserService.UpdateUser(id, payload)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondwithJSON(w, http.StatusOK, user)
}

// Delete user
func (h UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	// Check if the record exists
	err := h.UserService.DeleteUser(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
	}

	respondwithJSON(w, http.StatusNoContent, "User successfully deleted")
}

// User questions
func (h UserHandler) GetUserQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	user, err := h.UserService.GetUserQuestions(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
	}

	respondwithJSON(w, http.StatusOK, user)
}

// User answers
func (h UserHandler) GetUserAnswersHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	user, err := h.UserService.GetUserAnswers(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
	}

	respondwithJSON(w, http.StatusOK, user)
}
