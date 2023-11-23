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

type TagHandler struct {
	TagService service.TagService
}

func NewTagHandler(service service.TagService) *TagHandler {
	return &TagHandler{
		TagService: service,
	}
}

// All Tags
func (h TagHandler) GetAllTagsHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	itemsPerPage, _ := strconv.Atoi(r.URL.Query().Get("itemsPerPage"))
	if r.URL.Query().Get("page") == "" || r.URL.Query().Get("itemsPerPage") == "" {
		page, itemsPerPage = 1, 15
	}
	tags, err := h.TagService.GetAllTags(page, itemsPerPage)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{"data": tags, "page": page})
}

// Single Tag
func (h TagHandler) GetTagHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	tag, err := h.TagService.GetTag(id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, tag)
}

// Update Tag
func (h TagHandler) UpdateTagHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload := new(model.UpdateTagInput)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	tag, err := h.TagService.UpdateTag(id, payload)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, tag)
}

// Delete Tag
func (h TagHandler) DeleteTagHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	// Check if the record exists
	err := h.TagService.DeleteTag(id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusNoContent, "Tag successfully deleted")
}

// Tag Questions
func (h TagHandler) GetTagQuestionHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	tag, err := h.TagService.GetTagQuestions(id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, tag)
}
