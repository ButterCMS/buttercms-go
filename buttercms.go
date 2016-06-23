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
	url := API_ROOT_URL + path + "?auth_token=" + authToken

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
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

func GetPosts() (*PostAPIResponse, error) {
	body, err := getRequest("posts")

	var resp = new(PostAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
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
