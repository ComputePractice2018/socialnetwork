package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ComputePractice2018/socialnetwork/backend/data"
	"github.com/gorilla/mux"
)

//
//GetUsers обрабатывает запросы на получение списка пользователей
func GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(data.GetUsers())
	if err != nil {
		message := fmt.Sprintf("Unable to encode data: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
		return
	}

	//_, err = w.Write(binaryData)
	//if err != nil {
	//	message := fmt.Sprintf("Handler write error: %v", err)
	//	http.Error(w, message, http.StatusInternalServerError)
	//	log.Println(message)
	//}

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
	id := data.AddUser(user)
	w.Header().Add("Location", r.URL.String()+"/"+stronv.Itoa(id))
	w.WriteHeader(http.StatusCreated)
}

// EditUser обрабатывает PUT запрос
func EditUser(w http.ResponseWriter, r *http.Request) {
	var user data.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		message := fmt.Sprintf("Unable to decode PUT data: %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Println(message)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		message := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}

	err = data.EditUser(user, id)
	if err != nil {
		message := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}
	w.Header().Add("Location", r.URL.String())
	w.WriteHeader(http.StatusAccepted)

}

//DeleteUser обрабатывает DELETE запрос
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		message := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}

	err = data.RemoveUser(user, id)
	if err != nil {
		message := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}
	w.Header().Add("Location", r.URL.String())
	w.WriteHeader(http.StatusNoContent)

}
