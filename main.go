package main

import (
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

	r :=mux.NewRouter()


	r.HandleFunc("/",homeHandler).Methods("GET")
	r.HandleFunc("/api/hello", helloHandler).Methods("GET")

	fmt.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080",r))

}

func homeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the home page!")

}
func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello from Mux");
}