package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sinisaos/chi-ent/pkg/model"
	"github.com/sinisaos/chi-ent/pkg/service"
	"github.com/sinisaos/chi-ent/pkg/utils"

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
	if r.URL.Query().Get("page") == "" || r.URL.Query().Get("itemsPerPage") == "" {
		page, itemsPerPage = 1, 15
	}
	users, err := h.UserService.GetAllUsers(page, itemsPerPage)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{"data": users, "page": page})
}

// Single user
func (h UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	user, err := h.UserService.GetUser(id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, user)
}

// New user
func (h UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	payload := new(model.NewUserInput)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.UserService.CreateUser(payload)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, user)
}

// Update user
func (h UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	payload := new(model.UpdateUserInput)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	user, err := h.UserService.UpdateUser(id, payload)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, user)
}

// Delete user
func (h UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	// Check if the record exists
	err := h.UserService.DeleteUser(id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusNoContent, "User successfully deleted")
}

// User questions
func (h UserHandler) GetUserQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	user, err := h.UserService.GetUserQuestions(id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, user)
}

// User answers
func (h UserHandler) GetUserAnswersHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	user, err := h.UserService.GetUserAnswers(id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, user)
}
