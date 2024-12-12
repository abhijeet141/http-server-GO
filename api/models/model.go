package models

type User struct {
	Name         string
	HealthStatus []string
}

type KidneyRequest struct {
	NewKidney string `json:"newKidney"`
}

type KidneyDeleteRequest struct {
	DeleteKidney string `json:"deleteKidney"`
}
