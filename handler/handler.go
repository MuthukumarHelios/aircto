package handler

import (
	"aircto/dal"
	"aircto/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

type HandlerInstance struct {
}

func (h *HandlerInstance) RegisterUser(w http.ResponseWriter, r *http.Request) {

	userDal := &dal.Dal{}
	userData := models.User{}

	err := json.NewDecoder(r.Body).Decode(&userData)

	if err != nil {
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

func (h *HandlerInstance) Login(w http.ResponseWriter, r *http.Request) {

	reponseData := map[string]interface{}{
		// "message":"user Registed"
	}

	var payload map[string]string
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		With200(w, err.Error())
		return
	}

	tkn, err := CreateToken(payload["userName"])
	reponseData["token"] = tkn
	if err != nil {
		With200(w, err.Error())
		return
	}
	With200(w, reponseData)
	return
}

func (h *HandlerInstance) CreateIssue(w http.ResponseWriter, r *http.Request) {

	var payload map[string]string
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		With200(w, err.Error())
		return
	}
	fmt.Println("--", payload["token"])
	err := VerfifyToken(payload["userName"], payload["token"])

	if err != nil {
		With200(w, err.Error())
		return
	}
	With200(w, "You have successfully Created the Issue")
	return
}

func CreateToken(userName string) (string, error) {

	claims := map[string]interface{}{
		"userName": userName,
	}
	jwt := jws.NewJWT(claims, crypto.SigningMethodHS256)

	jwtToken, err := jwt.Serialize([]byte(nil))

	if err != nil {
		return "", err
	}
	return string(jwtToken), nil
}

func VerfifyToken(userName, token string) error {
	parsedToken, err := jws.ParseJWT([]byte(string(token)))
	if err != nil {
		return err
	}

	tokenUserName := parsedToken.Claims()["userName"]

	if tokenUserName != userName {
		return errors.New("You are not a valid user")
	}
	// sessionSecret := 123
	err = (parsedToken.Validate([]byte(nil), crypto.SigningMethodHS256))
	return err
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
