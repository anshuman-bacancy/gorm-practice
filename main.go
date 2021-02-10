package main

import (
	"net/http"
	"fake.com/anshuman/server"
	"fake.com/anshuman/handler"

	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.HomeHandler).Methods("GET")
	router.HandleFunc("/savebook", handler.AddBookHandler).Methods("POST")
	router.HandleFunc("/getbook/{name}", handler.GetBookHandler).Methods("GET")
	router.HandleFunc("/updatebook/{name}", handler.UpdateBookHandler).Methods("PUT")
	router.HandleFunc("/deletebook/{name}", handler.DeleteBookHandler).Methods("DELETE")

	server.InitializeDatabase()
	http.ListenAndServe(":8080", router)
}

