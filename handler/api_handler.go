package handler

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	respondwithJSON(w, http.StatusOK, map[string]interface{}{"message": "Echo ENT web api"})
}
