package main

import (
	"CodeChef_SIESGST_User_Tracker/controllers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	var err = godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	fmt.Println("Server Starting")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	}).Methods("GET")
	router.HandleFunc("/track", controllers.Track).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
