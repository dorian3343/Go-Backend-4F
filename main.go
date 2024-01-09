package main

import (
	"backend/handler"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Define routes using Gorilla Mux
	r.HandleFunc("/api/createUser", handler.UserCreation)
	r.HandleFunc("/api/deleteUser", handler.UserDeletion)
	r.HandleFunc("/api/addToBasket", handler.HandleBasketAdd)
	r.HandleFunc("/api/loginUser", handler.HandleUserLogin)
	r.HandleFunc("/api/getProducts", handler.HandleProductRequest)
	r.HandleFunc("/api/removeBaskets", handler.HandleBasketRemove)
	// CORS config
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// Use the CORS middleware
	http.Handle("/", handlers.CORS(headers, methods, origins)(r))

	// Starting HTTP server
	fmt.Println("Starting server on port:80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
