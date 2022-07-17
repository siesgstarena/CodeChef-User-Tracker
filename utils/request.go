package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func MakeGetRequest(url string, headers http.Header) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header = headers

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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}
