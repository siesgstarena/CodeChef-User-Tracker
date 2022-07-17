package helper

import (
	"CodeChef_SIESGST_User_Tracker/types"
	"fmt"
	"strings"
	"time"
)

func ExtractContestCodes(contestList []types.Contest, startDate string, endDate string) []string {
	var contestCodes []string
	var dateFormat = "2006-01-02"
	var startDateActual time.Time
	var endDateActual time.Time
	startDateActual, err := time.Parse(dateFormat, strings.TrimSpace(startDate))
	if err != nil {
		fmt.Println("Error In Parsing Start Date", err)
	}
	endDateActual, err = time.Parse(dateFormat, strings.TrimSpace(endDate))
	if err != nil {
		fmt.Println("Error in parsing end date", err)
	}
	for i, contest := range contestList {
		contestStartDate, err := time.Parse(dateFormat, strings.TrimSpace(contest.StartDate[0:10]))
		if err != nil {
			fmt.Println(err)
		}
		contestEndDate, err := time.Parse(dateFormat, strings.TrimSpace(contest.EndDate[0:10]))
		if err != nil {
			fmt.Println(err)
		}
		if i < 4 {
			fmt.Println("contestStartDate", contestStartDate, " contestEndDate ", contestEndDate, "contestCode", contest.Code)
		}
		if (contestStartDate.After(startDateActual) || contestStartDate.Equal(startDateActual)) && (contestEndDate.Before(endDateActual) || contestEndDate.Equal(endDateActual)) {
			contestCodes = append(contestCodes, contest.Code)
		}
	}
	return contestCodes
}

func UserHasSolved(userInfo types.ProblemsCategory, contestCodes []string) []string {
	var HasParticipated []string
	PartiallySolved := userInfo.PartiallySolved
	Solved := userInfo.Solved
	Attempted := userInfo.Attempted
	for _, contestCode := range contestCodes {
		if _, ok := PartiallySolved.(map[string]interface{}); ok {
			if _, ok := PartiallySolved.(map[string]interface{})[contestCode]; ok {
				HasParticipated = append(HasParticipated, contestCode)
				continue
			}
		}
		if _, ok := Solved.(map[string]interface{}); ok {
			if _, ok := Solved.(map[string]interface{})[contestCode]; ok {
				HasParticipated = append(HasParticipated, contestCode)
				continue
			}
		}
		if _, ok := Attempted.(map[string]interface{}); ok {
			if _, ok := Attempted.(map[string]interface{})[contestCode]; ok {
				HasParticipated = append(HasParticipated, contestCode)
				continue
			}
		}
	}
	return HasParticipated
}
