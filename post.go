package ButterCMS

import (
	"time"
)

type Post struct {
	URL              string     `json:"url"`
	Created          time.Time  `json:"created"`
	Published        time.Time  `json:"published"`
	Author           Author     `json:"author"`
	CategoryList     []Category `json:"categories"`
	TagList          []Tag      `json:"tags"`
	FeaturedImage    string     `json:"featured_image"`
	FeaturedImageAlt string     `json:"featured_image_alt"`
	Slug             string     `json:"slug"`
	Title            string     `json:"title"`
	Body             string     `json:"body"`
	Summary          string     `json:"summary"`
	SEOTitle         string     `json:"seo_title"`
	MetaDescription  string     `json:"meta_description"`
	Status           string     `json:"status"`
}
