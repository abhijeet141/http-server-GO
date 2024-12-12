package routes

import (
	"server/api/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.GetHandler).Methods("GET")
	r.HandleFunc("/", handlers.PostHandler).Methods("POST")
	r.HandleFunc("/", handlers.DeleteHandler).Methods("DELETE")
	return r
}
