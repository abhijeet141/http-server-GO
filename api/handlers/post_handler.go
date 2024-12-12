package handlers

import (
	"encoding/json"
	"net/http"
	"server/api/models"
	"server/api/services"
	"strings"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	var req models.KidneyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	if user == "" {
		http.Error(w, "User Query Parameter Required", http.StatusBadRequest)
		return
	}
	if req.NewKidney != "healthy" && req.NewKidney != "unhealthy" {
		http.Error(w, "Kidney must be healthy or unhealthy", http.StatusForbidden)
		return
	}
	patient := services.FindUserByName(userList, user)
	if patient == nil {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}
	if len(patient.HealthStatus) >= 2 {
		http.Error(w, "Cannot add more than 2 kidneys", http.StatusBadRequest)
		return
	}
	services.AddKidney(patient, strings.ToLower(req.NewKidney))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"Message": "Kidney added"})
}
