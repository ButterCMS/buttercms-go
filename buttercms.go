package ButterCMS

import (
	"strings"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var authToken string

const (
	API_ROOT_URL = "https://api.buttercms.com/v2/"
)

func getRequest(path string, params map[string]string) ([]byte, error) {
	if "" == authToken {
		return nil, errors.New("No auth token set")
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", API_ROOT_URL+path, nil)
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("auth_token", authToken)
	for k := range params {
		q.Add(k, params[k])
	}
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		return nil, errors.New(http.StatusText(resp.StatusCode))
	}

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

func GetFeed(feedType string) (*FeedAPIResponse, error) {
	body, err := getRequest("feeds/"+feedType, nil)
	if err != nil {
		return nil, err
	}

	var resp = new(FeedAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func SearchPosts(query string, params map[string]string) (*PostsAPIResponse, error) {
	params["query"] = query
	body, err := getRequest("search", params)
	if err != nil {
		return nil, err
	}

	var resp = new(PostsAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetPosts(params map[string]string) (*PostsAPIResponse, error) {
	body, err := getRequest("posts", params)
	if err != nil {
		return nil, err
	}

	var resp = new(PostsAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetPost(slug string) (*PostAPIResponse, error) {
	body, err := getRequest("posts/"+slug, nil)
	if err != nil {
		return nil, err
	}

	var resp = new(PostAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetAuthors(params map[string]string) (*AuthorsAPIResponse, error) {
	body, err := getRequest("authors", params)
	if err != nil {
		return nil, err
	}

	var resp = new(AuthorsAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetAuthor(slug string, params map[string]string) (*AuthorAPIResponse, error) {
	body, err := getRequest("authors/"+slug, params)
	if err != nil {
		return nil, err
	}

	var resp = new(AuthorAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetCategories(params map[string]string) (*CategoriesAPIResponse, error) {
	body, err := getRequest("categories", params)
	if err != nil {
		return nil, err
	}

	var resp = new(CategoriesAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetCategory(slug string, params map[string]string) (*CategoryAPIResponse, error) {
	body, err := getRequest("categories/"+slug, params)
	if err != nil {
		return nil, err
	}

	var resp = new(CategoryAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetContentFields(keys []string) (*ContentFieldsAPIResponse, error) {
	params := map[string]string{"keys": strings.Join(keys, ",")}
	body, err := getRequest("content", params)
	if err != nil {
		return nil, err
	}

	var resp = new(ContentFieldsAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
