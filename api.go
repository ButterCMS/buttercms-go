package ButterCMS

type PostsMeta struct {
	Count        int `json:"count"`
	NextPage     int `json:"next_page"`
	PreviousPage int `json:"previous_page"`
}

type PostsAPIResponse struct {
	MetaData PostsMeta `json:"meta"`
	PostList []Post    `json:"data"`
}

type PostMeta struct {
	NextPost     Post `json:"next_post"`
	PreviousPost Post `json:"previous_post"`
}

type PostAPIResponse struct {
	MetaData PostMeta `json:"meta"`
	Post     Post     `json:"data"`
}

type CategoryAPIResponse struct {
	CategoryList []Category `json:"data"`
}

type AuthorAPIResponse struct {
	AuthorList []Author `json:"data"`
}

type FeedAPIResponse struct {
	Data string `json:"data"`
}
