package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sinisaos/chi-ent/model"
	"github.com/sinisaos/chi-ent/service"
	"github.com/sinisaos/chi-ent/utils"

	"github.com/go-chi/chi/v5"
)

type QuestionHandler struct {
	QuestionService service.QuestionService
}

func NewQuestionHandler(service service.QuestionService) *QuestionHandler {
	return &QuestionHandler{
		QuestionService: service,
	}
}

// All Questions
func (h QuestionHandler) GetAllQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	itemsPerPage, _ := strconv.Atoi(r.URL.Query().Get("itemsPerPage"))
	if r.URL.Query().Get("page") == "" || r.URL.Query().Get("itemsPerPage") == "" {
		page, itemsPerPage = 1, 15
	}
	questions, err := h.QuestionService.GetAllQuestions(page, itemsPerPage)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{"data": questions, "page": page})
}

// Single Question
func (h QuestionHandler) GetQuestionHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	question, err := h.QuestionService.GetQuestion(id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, question)
}

// New Question
func (h QuestionHandler) CreateQuestionHandler(w http.ResponseWriter, r *http.Request) {
	payload := new(model.NewQuestionInput)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	question, err := h.QuestionService.CreateQuestion(payload)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, question)
}

// Update Question
func (h QuestionHandler) UpdateQuestionHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload := new(model.UpdateQuestionInput)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	question, err := h.QuestionService.UpdateQuestion(id, payload)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, question)
}

// Delete Question
func (h QuestionHandler) DeleteQuestionHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	// Check if the record exists
	err := h.QuestionService.DeleteQuestion(id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusNoContent, "Question successfully deleted")
}

// Question Answers
func (h QuestionHandler) GetQuestionAnswersHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	question, err := h.QuestionService.GetQuestionAnswers(id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, question)
}

// Question Author
func (h QuestionHandler) GetQuestionAuthorHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	question, err := h.QuestionService.GetQuestionAuthor(id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, question)
}

// Question Tags
func (h QuestionHandler) GetQuestionTagsHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	question, err := h.QuestionService.GetQuestionTags(id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, question)
}
