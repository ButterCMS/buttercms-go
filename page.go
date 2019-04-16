package ButterCMS

type Page struct {
	Slug 		string                  `json:"slug"`
	PageType 	string                  `json:"page_type"`
	Fields 		map[string]interface{}	`json:"fields"`
}
