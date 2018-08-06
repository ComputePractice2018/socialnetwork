package main

import (
	"log"
	"net/http"

	"github.com/ComputePractice2018/socialnetwork/backend/data"
)


func main() {
	
	http.HandleFunc("/api/socialnetwork/users", server.GetUsers) {

	log.Fatal(http.ListenAndServe(":8080", nil))
}
