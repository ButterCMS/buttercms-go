package ButterCMS

type PagingMeta struct {
	Count        int `json:"count"`
	NextPage     int `json:"next_page"`
	PreviousPage int `json:"previous_page"`
}

// Posts
type PostsAPIResponse struct {
	MetaData PagingMeta `json:"meta"`
	PostList []Post     `json:"data"`
}

type PostMeta struct {
	NextPost     Post `json:"next_post"`
	PreviousPost Post `json:"previous_post"`
}

type PostAPIResponse struct {
	MetaData PostMeta	`json:"meta"`
	Post     Post		`json:"data"`
}

// Pages
type PagesMeta struct {
	Count        int `json:"count"`
	NextPage     int `json:"next_page"`
	PreviousPage int `json:"previous_page"`
}

type PagesAPIResponse struct {
	MetaData PagingMeta	`json:"meta"`
	PageList []Page		`json:"data"`
}

type PageAPIResponse struct {
	Page Page `json:"data"`
}

// Categories
type CategoriesAPIResponse struct {
	CategoryList []Category `json:"data"`
}

type CategoryAPIResponse struct {
	Category Category `json:"data"`
}

// Tags
type TagsAPIResponse struct {
	TagList []Tag `json:"data"`
}

type TagAPIResponse struct {
	Tag Tag `json:"data"`
}

// Authors
type AuthorsAPIResponse struct {
	AuthorList []Author `json:"data"`
}

type AuthorAPIResponse struct {
	Author Author `json:"data"`
}

// Feed
type FeedAPIResponse struct {
	Data string `json:"data"`
}

// Content Fields
type ContentFieldsAPIResponse struct {
	Data map[string]string `json:"data"`
}
