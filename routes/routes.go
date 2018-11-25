package routes

import (
  "github.com/gorilla/mux"
  "aircto/handler"
)



func MainRoutes()*mux.Router{
	
 	h := &handler.HandlerInstance{}

	r := mux.NewRouter()
	
	r.HandleFunc("/api/user/",h.RegisterUser).Methods("POST")	



	return r
}