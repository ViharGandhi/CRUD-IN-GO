package main

import (
	"fmt"
	"log"
	"net/http"
	"todo/connections"
	"todo/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to MongoDB
	connections.Connection()

	// Initialize Router
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/todo", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/todo", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/todo/{id}", handlers.GetTodoByID).Methods("GET")
	r.HandleFunc("/todo/{id}", handlers.UpdateTodoByID).Methods("PUT")
	r.HandleFunc("/todo/{id}", handlers.DeleteTodoByID).Methods("DELETE")

	// Start the server
	port := "3001"
	fmt.Println("Server running at port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
