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

type AnswerHandler struct {
	AnswerService service.AnswerService
}

func NewAnswerHandler(service service.AnswerService) *AnswerHandler {
	return &AnswerHandler{
		AnswerService: service,
	}
}

// All Answers
func (h AnswerHandler) GetAllAnswersHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	itemsPerPage, _ := strconv.Atoi(r.URL.Query().Get("itemsPerPage"))
	if r.URL.Query().Get("page") == "" || r.URL.Query().Get("itemsPerPage") == "" {
		page, itemsPerPage = 1, 15
	}
	answers, err := h.AnswerService.GetAllAnswers(page, itemsPerPage)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{"data": answers, "page": page})
}

// Single Answer
func (h AnswerHandler) GetAnswerHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	answer, err := h.AnswerService.GetAnswer(id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, answer)
}

// New Answer
func (h AnswerHandler) CreateAnswerHandler(w http.ResponseWriter, r *http.Request) {
	payload := new(model.NewAnswerInput)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	answer, err := h.AnswerService.CreateAnswer(payload)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, answer)
}

// Update Answer
func (h AnswerHandler) UpdateAnswerHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload := new(model.UpdateAnswerInput)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	answer, err := h.AnswerService.UpdateAnswer(id, payload)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, answer)
}

// Delete Answer
func (h AnswerHandler) DeleteAnswerHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	// Check if the record exists
	err := h.AnswerService.DeleteAnswer(id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusNoContent, "Answer successfully deleted")
}

// Answer Questions
func (h AnswerHandler) GetAnswerQuestionHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	answer, err := h.AnswerService.GetAnswerQuestion(id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, answer)
}

// Answer Author
func (h AnswerHandler) GetAnswerAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	answer, err := h.AnswerService.GetAnswerAuthor(id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, answer)
}
