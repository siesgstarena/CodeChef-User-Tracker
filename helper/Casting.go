package helper

import (
	"CodeChef_SIESGST_User_Tracker/types"
	"encoding/json"
	"fmt"
)

func ConvertToContestInfo(body []byte) ([]types.Contest, error) {
	var users types.ContestInfo
	err := json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users.Result.Data.Content.ContestList, nil

}
func ConvertToUserInfo(body []byte) (types.UserInfo, error) {
	var users types.UserInfo
	err := json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println(err)
		return users, err
	}
	return users, nil
}
