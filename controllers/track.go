package controllers

import (
	"CodeChef_SIESGST_User_Tracker/helper"
	"CodeChef_SIESGST_User_Tracker/public"
	"CodeChef_SIESGST_User_Tracker/services"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gorilla/schema"
	// "github.com/gorilla/websocket"
)

type FormElement struct {
	Startingdate string `json:"startingdate"`

	Endingdate string `json:"endingdate"`

	Datafile *multipart.FileHeader `json:"datafile"`
}

func Track(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Println("Error in parsing form", err)
	}
	var form FormElement
	decoder := schema.NewDecoder()
	err = decoder.Decode(&form, r.PostForm)
	if err != nil {
		fmt.Println("Error in decoding form", err)
	}
	bookpath, err := helper.SaveFileToDestination(r)
	if err != nil {
		fmt.Println("Error in saving file", err)
	}
	contestCodes, err := services.CodeChefContest(form.Startingdate, form.Endingdate)
	fmt.Println("contestCodes ", contestCodes)
	if err != nil {
		fmt.Println("Error in making request", err)
	}
	usernames, err := helper.ConvertExcellToArray(bookpath)
	if err != nil {
		fmt.Println("Error in making request", err)
	}
	usersSolved, err := public.UsersHaveSolved(usernames, contestCodes)
	if err != nil {
		fmt.Println("Error in making request", err)
	}
	err = helper.WriteToExcell(bookpath, usersSolved, form.Startingdate, form.Endingdate)
	if err != nil {
		fmt.Println("Error in making request", err)
	}
	fmt.Fprintf(w, "Hello World")
}
