package main

import (
	"log"
	"net/http"

	"github.com/DSY4444/go-task-manager/auth"
	"github.com/DSY4444/go-task-manager/database"
	"github.com/DSY4444/go-task-manager/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	database.ConnectDB()

	r := mux.NewRouter()

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		token, _ := auth.GenerateToken()
		w.Write([]byte(token))
	}).Methods("POST")

	// r.HandleFunc("/tasks", auth.AuthMiddleware(handlers.GetTasks)).Methods("GET")
	// r.HandleFunc("/tasks", auth.AuthMiddleware(handlers.CreateTask)).Methods("POST")
	// r.HandleFunc("/tasks/{id}", auth.AuthMiddleware(handlers.UpdateTask)).Methods("PUT")
	// r.HandleFunc("/tasks/{id}", auth.AuthMiddleware(handlers.DeleteTask)).Methods("DELETE")

	// Enable CORS
	corsHandler := cors.New(cors.Options{
		// AllowedOrigins:   []string{"http://localhost:3000"}, // Change to your frontend URL in production
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	log.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", corsHandler.Handler(r))
}
