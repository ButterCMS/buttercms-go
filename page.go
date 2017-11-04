package ButterCMS

type Page struct {
	Slug 	string                  `json:"slug"`
	Fields 	map[string]interface{}	`json:"fields"`
}
