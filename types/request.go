package types

type ContestInfo struct {
	Status string `json:"status"`
	Result struct {
		Data struct {
			Content struct {
				CurrentTime int       `json:"currentTime"`
				ContestList []Contest `json:"contestList"`
			} `json:"content"`
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"data"`
	} `json:"result"`
}
type Contest struct {
	Code      string `json:"code"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type UserInfo struct {
	Status string `json:"status"`
	Result struct {
		Data struct {
			Content struct {
				Username     string           `json:"username"`
				Fullname     string           `json:"fullname"`
				ProblemStats ProblemsCategory `json:"problemStats"`
			} `json:"content"`
		} `json:"data"`
	} `json:"result"`
}

type ProblemsCategory struct {
	PartiallySolved interface {
	} `json:"partiallySolved"`
	Solved interface {
	} `json:"solved"`
	Attempted interface {
	} `json:"attempted"`
}

type Token struct {
	Status string `json:"status"`
	Result struct {
		Data struct {
			Access_token string `json:"access_token"`
			Expires_in   int    `json:"expires_in"`
			Token_type   string `json:"token_type"`
			Scope        string `json:"scope"`
		} `json:"data"`
	} `json:"result"`
}
