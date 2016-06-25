package ButterCMS

type Category struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	RecentPosts []Post `json:"recent_posts"`
}
