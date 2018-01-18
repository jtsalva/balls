package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		return
	}

	var loginRequest LoginRequest
	var loginResponse LoginResponse

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		log.Println(err)
		return
	}

	log.Println(loginRequest.Name)

	loginResponse.Id = 1234
	loginResponse.Name = loginRequest.Name

	if loginRequest.Password == "thispassword" {
		loginResponse.Success = true
	} else {
		loginResponse.Success = false
	}

	if err := json.NewEncoder(w).Encode(loginResponse); err != nil {
		log.Println(err)
		return
	}

}
