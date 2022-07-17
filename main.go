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
	router.HandleFunc("/track", controllers.Track)
	router.HandleFunc("/", controllers.Home).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
