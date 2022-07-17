package controllers

import (
	"CodeChef_SIESGST_User_Tracker/helper"
	"CodeChef_SIESGST_User_Tracker/public"
	"CodeChef_SIESGST_User_Tracker/services"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

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
		fmt.Fprintf(w, "Error in parsing form")
	}
	var form FormElement
	decoder := schema.NewDecoder()
	err = decoder.Decode(&form, r.PostForm)
	if err != nil {
		fmt.Println("Error in decoding form", err)
		fmt.Fprintf(w, "Error in decoding form")
	}
	filepath, err := helper.SaveFileToDestination(r)
	if err != nil {
		fmt.Println("Error in saving file", err)
		fmt.Fprintf(w, "Error in saving file")
	}
	contestCodes, err := services.CodeChefContest(form.Startingdate, form.Endingdate)
	fmt.Println("contestCodes ", contestCodes)
	if err != nil {
		fmt.Println("Error in making request", err)
		fmt.Fprintf(w, "Error in making request")
	}
	usernames, err := helper.ConvertExcellToArray(filepath)
	if err != nil {
		fmt.Println("Error in making request", err)
		fmt.Fprintf(w, "Error in making request")
	}
	usersSolved, err := public.UsersHaveSolved(usernames, contestCodes)
	if err != nil {
		fmt.Println("Error in making request", err)
		fmt.Fprintf(w, "Error in making request")
	}
	filename, err := helper.WriteToExcell(filepath, usersSolved, form.Startingdate, form.Endingdate)
	if err != nil {
		fmt.Println("Error in making request", err)
		fmt.Fprintf(w, "Error in writing to excell")
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filepath))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, filename)
}
