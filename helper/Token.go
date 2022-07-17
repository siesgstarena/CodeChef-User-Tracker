package helper

import (
	"CodeChef_SIESGST_User_Tracker/types"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var Token = ""
var Expires_at = time.Now().UTC()

func MakePostRequest(url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("content-Type", "application/json")
	if err != nil {
		fmt.Println("Error in making request", err)
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error in making request", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}
func MakeToken() string {
	fmt.Println("Token", Token)
	if len(Token) == 0 || Expires_at.Before(time.Now().UTC()) {
		fmt.Println("f")
		url := "https://api.codechef.com/oauth/token"
		payload := map[string]string{
			"grant_type":    "client_credentials",
			"scope":         "public",
			"client_id":     os.Getenv("CLIENT_ID"),
			"client_secret": os.Getenv("CLIENT_SECRET"),
		}
		// map[string]string to []byte
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		resp, err := MakePostRequest(url, payloadBytes)
		if err != nil {
			fmt.Println("Error in making request", err)
		}
		var token types.Token
		err = json.Unmarshal(resp, &token)
		if err != nil {
			fmt.Println("Error in making request", err)
		}
		Token = token.Result.Data.Access_token
		Expires_at = time.Now().UTC().Add(time.Duration(token.Result.Data.Expires_in) * time.Second)
	}
	return Token
}
