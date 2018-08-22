package server

//
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-Type", "text/plain;charset=utf-8")
		fmt.Fprint(w, "JSON marshaling error: %v", err)
		return
	}

	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(binaryData)
	if err != nil {
		log.Printf("Handler write error: %v", err)
	}
}
