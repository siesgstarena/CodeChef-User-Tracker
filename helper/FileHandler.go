package helper

import (
	"CodeChef_SIESGST_User_Tracker/types"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func SaveFileToDestination(r *http.Request) (string, error) {
	_, file, err := r.FormFile("datafile")
	if err != nil {
		fmt.Println("Error in getting file", err)
		return "", err
	}
	src, err := file.Open()
	if err != nil {
		fmt.Println("Error in opening file", err)
		return "", err
	}
	defer src.Close()
	name, _ := os.Getwd()
	fileextension := filepath.Ext(file.Filename)
	newfilename := filepath.Base("datas" + fileextension)
	newfilepath := filepath.Join(name, newfilename)

	dst, err := os.Create(newfilepath)
	if err != nil {
		fmt.Println("Error in creating file", err)
		return "", err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println("Error in copying file", err)
		return "", err
	}
	if err != nil {
		fmt.Println("Error in getting bookfile", err)
		return "", err
	}
	return newfilename, nil
}
func ConvertExcellToArray(file string) ([]string, error) {
	var data []string
	content, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println("Error in opening file", err)
		return data, err
	}
	cols, _ := content.GetCols("Sheet1")
	for i := 1; i < len(cols[0]); i++ {
		if cols[0][i] != "" {
			data = append(data, cols[1][i])
		}
	}
	// 	CodeChef id must be Second column
	// data = append(data, cols[1][1:]...)
	defer content.Close()
	return data, nil
}
func WriteToExcell(file string, usersHaveSolved []types.UserSolveds, startDate string, endDate string) error {
	content, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println("Error in opening file", err)
		return err
	}
	content.SetCellValue("Sheet1", "B1", "CodeChef")
	content.SetCellValue("Sheet1", "C1", "Participation "+startDate+" to "+endDate)
	content.SetCellValue("Sheet1", "D1", "Total Number Participation")
	for i, userSolved := range usersHaveSolved {
		fmt.Println("userSolved ", userSolved, strconv.Itoa((i + 2)))
		content.SetCellValue("Sheet1", "B"+strconv.Itoa((i+2)), userSolved.UserName)
		content.SetCellValue("Sheet1", "C"+strconv.Itoa((i+2)), userSolved.AllProblems)
		content.SetCellValue("Sheet1", "D"+strconv.Itoa((i+2)), len(userSolved.AllProblems))
	}
	// set the width of the column
	errs := content.SetColWidth("Sheet1", "C", "C", 100)
	if errs != nil {
		fmt.Println("Error in setting column width", err)
		return err
	}
	err = content.SetColWidth("Sheet1", "D", "D", 24)
	if err != nil {
		fmt.Println("Error in setting column width", err)
		return err
	}

	err = content.SaveAs("./output.xlsx")
	if err != nil {
		fmt.Println("Error in saving file", err)
		return err
	}
	defer content.Close()
	return nil
}
