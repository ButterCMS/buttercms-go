package ButterCMS

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var authToken string

var preview bool

const (
	VERSION      = "2.3.0"
	API_BASE_URL = "https://api.buttercms.com/"
	V2_ENDPOINT  = "/v2"
)

var ENV_BASE_URL = os.Getenv("API_BASE_URL")

func getRequest(path string, params map[string]string) ([]byte, error) {
	if authToken == "" {
		return nil, errors.New("no auth token set")
	}

	client := http.DefaultClient
	baseUrl := API_BASE_URL
	if ENV_BASE_URL != "" {
		baseUrl = ENV_BASE_URL
	}

	url, err := url.JoinPath(baseUrl, V2_ENDPOINT, path)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Butter-Client", "Go/"+VERSION)

	q := req.URL.Query()
	q.Add("auth_token", authToken)
	for k := range params {
		q.Add(k, params[k])
	}

	if preview {
		q.Add("preview", "1")
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	bodyBytes := []byte(body)

	return bodyBytes, err
}

func SetAuthToken(token string) {
	authToken = token
}

func SetPreviewMode(inPreview bool) {
	preview = inPreview
}

///////////
// Feed
///////////

func GetFeed(feedType string) (*FeedAPIResponse, error) {
	body, err := getRequest("feeds/"+feedType+"/", nil)
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

///////////
// Pages
///////////

func GetPages(pageType string, params map[string]string) (*PagesAPIResponse, error) {
	body, err := getRequest("pages/"+pageType+"/", params)
	if err != nil {
		return nil, err
	}

	var resp = new(PagesAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetPage(pageType string, slug string, params map[string]string) (*PageAPIResponse, error) {
	body, err := getRequest("pages/"+pageType+"/"+slug+"/", params)
	if err != nil {
		return nil, err
	}

	var resp = new(PageAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

///////////
// Posts
///////////

func SearchPosts(query string, params map[string]string) (*PostsAPIResponse, error) {
	params["query"] = query
	body, err := getRequest("search/", params)
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
	body, err := getRequest("posts/", params)
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
	body, err := getRequest("posts/"+slug+"/", nil)
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

///////////
// Authors
///////////

func GetAuthors(params map[string]string) (*AuthorsAPIResponse, error) {
	body, err := getRequest("authors/", params)
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
	body, err := getRequest("authors/"+slug+"/", params)
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

///////////
// Categories
///////////

func GetCategories(params map[string]string) (*CategoriesAPIResponse, error) {
	body, err := getRequest("categories/", params)
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

///////////
// Tags
///////////

func GetTags(params map[string]string) (*TagsAPIResponse, error) {
	body, err := getRequest("tags/", params)
	if err != nil {
		return nil, err
	}

	var resp = new(TagsAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetTag(slug string, params map[string]string) (*TagAPIResponse, error) {
	body, err := getRequest("tags/"+slug+"/", params)
	if err != nil {
		return nil, err
	}

	var resp = new(TagAPIResponse)
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

///////////
// Content Fields
///////////

func GetContentFields(keys []string, params map[string]string) (*ContentFieldsAPIResponse, error) {
	params["keys"] = strings.Join(keys, ",")
	body, err := getRequest("content/", params)
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
