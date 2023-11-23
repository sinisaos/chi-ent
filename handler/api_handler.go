package handler

import (
	"net/http"

	"github.com/sinisaos/chi-ent/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{"message": "Chi ENT web api"})
}
