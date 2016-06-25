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

func getRequest(path string, params map[string]string) ([]byte, error) {
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

	var resp = new(FeedAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func SearchPosts(query string, page int) (*PostsAPIResponse, error) {
	params := map[string]string{"query": query, "page": string(page)}
	body, err := getRequest("search", params)

	var resp = new(PostsAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetPosts() (*PostsAPIResponse, error) {
	body, err := getRequest("posts", nil)

	var resp = new(PostsAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetPost(slug string) (*PostAPIResponse, error) {
	body, err := getRequest("posts/"+slug, nil)

	var resp = new(PostAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetAuthors(includeRecentPosts bool) (*AuthorsAPIResponse, error) {
	params := make(map[string]string)
	if includeRecentPosts {
		params = map[string]string{"include": "recent_posts"}
	}
	body, err := getRequest("authors", params)

	var resp = new(AuthorsAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetAuthor(slug string) (*AuthorAPIResponse, error) {
	body, err := getRequest("authors/"+slug, nil)

	var resp = new(AuthorAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetCategories(includeRecentPosts bool) (*CategoriesAPIResponse, error) {
	params := make(map[string]string)
	if includeRecentPosts {
		params = map[string]string{"include": "recent_posts"}
	}
	body, err := getRequest("categories", params)

	var resp = new(CategoriesAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetCategory(slug string) (*CategoryAPIResponse, error) {
	body, err := getRequest("categories/"+slug, nil)

	var resp = new(CategoryAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
