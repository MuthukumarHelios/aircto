package handler

import (
	
	"net/http"
	"fmt"
	"aircto/dal"
	"encoding/json"
	"aircto/models"
)

type HandlerInstance struct{
}

func (h *HandlerInstance)RegisterUser(w http.ResponseWriter, r *http.Request){
	
	userDal := &dal.Dal{}
   userData := models.User{}
                     
	err := json.NewDecoder(r.Body).Decode(&userData)
	
	if err != nil{
		fmt.Fprintf(w, err.Error())   
		return 
	}
	userDal.InsertUser(userData)

	reponseData := map[string]interface{}{
		// "message":"user Registed"
	}
	
	reponseData["message"] = "user Registed"
	With200(w, reponseData)
	return
}









func With200(w http.ResponseWriter, data interface{}) {
	dataB, err := json.Marshal(data)
	if err != nil {

		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dataB)
}
