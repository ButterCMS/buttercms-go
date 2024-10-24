package ButterCMS

import (
	"time"
)

type PostStatus string

const (
	PostStatusDraft     PostStatus = "draft"
	PostStatusPublished PostStatus = "published"
	PostStatusScheduled PostStatus = "scheduled"
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
	Status           PostStatus `json:"status"`
	Scheduled        time.Time  `json:"scheduled,omitempty"`
}
