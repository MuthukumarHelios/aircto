package routes

import (
	"aircto/handler"

	"github.com/gorilla/mux"
)

func MainRoutes() *mux.Router {

	h := &handler.HandlerInstance{}

	r := mux.NewRouter()

	r.HandleFunc("/api/user/", h.RegisterUser).Methods("POST")
	r.HandleFunc("/api/user/login/", h.Login).Methods("POST")
	r.HandleFunc("/api/user/issue/", h.CreateIssue).Methods("POST")

	return r
}
