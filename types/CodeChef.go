package types

type UserSolveds struct {
	UserName    string   `json:"username"`
	AllProblems []string `json:"allproblems"`
}
