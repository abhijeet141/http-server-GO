package services

import (
	"server/api/models"
	"strings"
)

func FindUserByName(userList []models.User, name string) *models.User {
	for i, user := range userList {
		if user.Name == name {
			return &userList[i]
		}
	}
	return nil
}

func OrganCount(patient *models.User) int {
	return len(patient.HealthStatus)
}

func StatusUpdate(patient *models.User) []string {
	var healthStatus []string
	for _, status := range patient.HealthStatus {
		healthStatus = append(healthStatus, strings.ToLower(status))
	}
	return healthStatus
}

func AddKidney(patient *models.User, kidney string) {
	patient.HealthStatus = append(patient.HealthStatus, kidney)
}

func DeleteKidney(patient *models.User, kidney string) {
	var updatedStatus []string
	for _, status := range patient.HealthStatus {
		if strings.ToLower(status) != kidney {
			updatedStatus = append(updatedStatus, status)
		}
	}
	patient.HealthStatus = updatedStatus
}
