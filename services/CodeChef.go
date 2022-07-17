package services

import (
	"CodeChef_SIESGST_User_Tracker/helper"
	"CodeChef_SIESGST_User_Tracker/utils"
	"fmt"
	"net/http"
)

var datas = [2]string{"yourdad01", "saahil468"}

func CodeChefContest(startDate string, endDate string) ([]string, error) {
	var contestCodes []string
	url := "https://api.codechef.com/contests?fields=startDate,endDate,code&sortBy=startDate"
	headers := http.Header{
		"Accept":        {"application/json"},
		"Authorization": {"Bearer " + helper.MakeToken()},
	}
	resp, err := utils.MakeGetRequest(url, headers)
	if err != nil {
		fmt.Println("Error in making request", err)
		return contestCodes, err
	}
	resps, err := helper.ConvertToContestInfo(resp)
	if err != nil {
		fmt.Println("Error in making request", err)
		return contestCodes, err
	}
	// fmt.Println(string(resp))
	contestCodes = helper.ExtractContestCodes(resps, startDate, endDate)

	return contestCodes, nil

}

func UserHasSolver(username string, contestcodes []string) ([]string, error) {
	HasParticipated := []string{}
	url := "https://api.codechef.com/users/" + username + "?fields=problemStats"
	headers := http.Header{
		"Accept":        {"application/json"},
		"Authorization": {"Bearer " + helper.MakeToken()},
	}
	resp, err := utils.MakeGetRequest(url, headers)
	if err != nil {
		fmt.Println("Error in making request", err)
		return HasParticipated, err
	}
	resps, err := helper.ConvertToUserInfo(resp)
	if err != nil {
		fmt.Println("Error in making request", err)
		return HasParticipated, err
	}
	HasParticipated = helper.UserHasSolved(resps.Result.Data.Content.ProblemStats, contestcodes)
	fmt.Println("test")
	return HasParticipated, nil

}
