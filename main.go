package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gunash-portfolio/go-backend-project/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database.Connect()

	r := mux.NewRouter()

	r.HandleFunc("/api/hello", helloHandler).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello from Mux",
	})
}
