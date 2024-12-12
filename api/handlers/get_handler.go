package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/api/models"
	"server/api/services"
)

var userList = []models.User{
	{Name: "Abhijeet", HealthStatus: []string{"healthy", "healthy"}},
	{Name: "Aditya", HealthStatus: []string{"healthy", "unhealthy"}},
	{Name: "Akash", HealthStatus: []string{"healthy"}},
	{Name: "Rohit", HealthStatus: []string{"unhealthy", "unhealthy"}},
	{Name: "Lakshya", HealthStatus: []string{"unhealthy"}},
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	if user == "" {
		http.Error(w, "User Query Parameter Required", http.StatusBadRequest)
		return
	}
	patient := services.FindUserByName(userList, user)
	if patient == nil {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}
	kidneyCount := services.OrganCount(patient)
	status := services.StatusUpdate(patient)
	response := fmt.Sprintf("%s has %d kidney(s) and health is %v", user, kidneyCount, status)
	data, err := json.MarshalIndent(map[string]interface{}{"User": response}, "", "\t")
	if err != nil {
		http.Error(w, "Error Marshalling Data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
