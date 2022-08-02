package main

import (
	"CodeChef_SIESGST_User_Tracker/controllers"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/track", controllers.Track).Methods("POST")
	muxWithMiddlewares := http.TimeoutHandler(router, time.Second*300, "Timeout!")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), muxWithMiddlewares))
}
