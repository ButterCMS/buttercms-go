package ButterCMS

type Category struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type CategoryAPIResponse struct {
	CategoryList []Category `json:"data"`
}
