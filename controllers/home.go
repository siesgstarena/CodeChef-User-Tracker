package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
	http.ServeFile(w, r, "./templates/index.html")
}
