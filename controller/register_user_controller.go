package controller

import (
	_ "api/dao"
	"api/model"
	"api/usecase"
	"encoding/json"
	"log"
	"net/http"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	switch r.Method {

	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return

	case http.MethodPost:
		var postData model.UserRegister
		err := json.NewDecoder(r.Body).Decode(&postData)
		if err != nil {
			log.Printf("fail: json.NewDecoder, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = usecase.RegisterUser(postData)
		if err != nil {
			log.Printf("fail: usecase.RegisterUser, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Id string `json:"id"`
		}{Id: postData.Id})

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func UpdateUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	switch r.Method {

	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return

	case http.MethodPost:
		var postData model.UserRegister
		err := json.NewDecoder(r.Body).Decode(&postData)
		if err != nil {
			log.Printf("fail: json.NewDecoder, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = usecase.UpdateUserProfile(postData)
		if err != nil {
			log.Printf("fail: usecase.UpdateUserProfile, %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Success bool `json:"success"`
		}{Success: true})

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	}
}
