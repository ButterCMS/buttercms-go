package ButterCMS

type Post struct {
	URL             string     `json:"url"`
	Author          Author     `json:"author"`
	CategoryList    []Category `json:"categories"`
	FeaturedImage   string     `json:"featured_image"`
	Slug            string     `json:"slug"`
	Title           string     `json:"title"`
	Body            string     `json:"body"`
	Summary         string     `json:"summary"`
	SEOTitle        string     `json:"seo_title"`
	MetaDescription string     `json:"meta_description"`
	Status          string     `json:"status"`
}
