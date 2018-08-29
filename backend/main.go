package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ComputePractice2018/socialnetwork/backend/server"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/socialnetwork/users", server.GetUsers).Methods("GET")
	router.HandleFunc("/api/socialnetwork/users", server.AddUsers).Methods("POST")
	router.HandleFunc("/api/socialnetwork/users/{id}", server.EditUsers).Methods("PUT")
	router.HandleFunc("/api/socialnetwork/users/{id}", server.DeleteUsers).Methods("Delete")

	//http.HandleFunc("/api/socialnetwork/users", server.UserHandler) {

	log.Fatal(http.ListenAndServe(":8080", router))
}
