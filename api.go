package ButterCMS

type APIMeta struct {
	Count        int `json:"count"`
	NextPage     int `json:"next_page"`
	PreviousPage int `json:"previous_page"`
}

type PostAPIResponse struct {
	MetaData APIMeta `json:"meta"`
	PostList []Post  `json:"data"`
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
