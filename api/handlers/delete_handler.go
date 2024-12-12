package handlers

import (
	"encoding/json"
	"net/http"
	"server/api/models"
	"server/api/services"
	"strings"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	var req models.KidneyDeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	if user == "" {
		http.Error(w, "User Query Parameter Required", http.StatusBadRequest)
		return
	}
	if req.DeleteKidney != "healthy" && req.DeleteKidney != "unhealthy" {
		http.Error(w, "Kidney must be healthy or unhealthy", http.StatusForbidden)
		return
	}
	patient := services.FindUserByName(userList, user)
	if patient == nil {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}
	services.DeleteKidney(patient, strings.ToLower(req.DeleteKidney))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"Message": "Kidney deleted"})
}
