package ButterCMS

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var authToken string

const (
	API_ROOT_URL = "https://api.buttercms.com/v2/"
)

func getRequest(path string) ([]byte, error) {
	resp, err := http.Get(API_ROOT_URL + path + "?auth_token=" + authToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	bodyBytes := []byte(body)

	return bodyBytes, err
}

func SetAuthToken(token string) {
	authToken = token
}

func GetAuthors() (*AuthorAPIResponse, error) {
	body, err := getRequest("authors")

	var resp = new(AuthorAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetCategories() (*CategoryAPIResponse, error) {
	body, err := getRequest("categories")

	var resp = new(CategoryAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
