package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ComputePractice2018/socialnetwork/backend/data"
)

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

func EditUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
