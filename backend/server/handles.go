package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ComputePractice2018/socialnetwork/backend/data"
)

//UserHandler обрабатывает все запросы /api/socialnetwork/user
func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		GetUsers(w, r)
		return
	}
	if r.Method == "POST" {
		AddUser(w, r)
		return
	}
	message := fmt.Sprintf("Method %s is not allowed", r.Method)
	http.Error(w, message, http.StatusMethodNotAllowed)
	log.Println(message)
}

//GetUsers обрабатывает запросы на получение списка пользователей
func GetUsers(w http.ResponseWriter, r *http.Request) {
	binaryData, err := json.Marshal(data.UserList)

	if err != nil {
		message := fmt.Sprintf("JSON marshaling error: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
		return
	}

	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(binaryData)
	if err != nil {
		message := fmt.Sprintf("Handler write error: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
	}

}

//AddUser обрабатывает POST запрос
func AddUser(w http.ResponseWriter, r *http.Request) {
	var user data.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		message := fmt.Sprintf("Unable to decode POST data: %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Println(message)
		return
	}
	log.Printf("%+v", user)
	w.WriteHeader(http.StatusCreated)
}
