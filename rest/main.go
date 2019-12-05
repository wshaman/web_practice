package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wshaman/web_practice/rest/handlers/contact"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/contact", contact.HandleList).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/contact", contact.HandleCreate).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/contact/{id}", contact.HandleUpdate).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/contact/{id}", contact.HandleLoad).Methods(http.MethodGet)
	if err := http.ListenAndServe(":8083", router); err != nil {
		log.Fatal(err)
	}
}
