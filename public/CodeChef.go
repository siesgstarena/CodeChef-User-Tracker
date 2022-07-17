package public

import (
	"CodeChef_SIESGST_User_Tracker/services"
	"CodeChef_SIESGST_User_Tracker/types"
)

func UsersHaveSolved(usernames []string, contestcodes []string) ([]types.UserSolveds, error) {
	var usersSolved []types.UserSolveds
	for _, username := range usernames {
		var userSolved types.UserSolveds
		userSolved.UserName = username
		userSolved.AllProblems, _ = services.UserHasSolver(username, contestcodes)
		usersSolved = append(usersSolved, userSolved)
	}
	return usersSolved, nil
}
