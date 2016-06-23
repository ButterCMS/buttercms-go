package ButterCMS

type Author struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Slug          string `json:"slug"`
	Bio           string `json:"bio"`
	Title         string `json:"title"`
	LinkedInUrl   string `json:"linkedin_url"`
	FacebookUrl   string `json:"facebook_url"`
	PinterestUrl  string `json:"pinterest_url"`
	InstagramUrl  string `json:"instagram_url"`
	TwitterHandle string `json:"twitter_handle"`
	ProfileImage  string `json:"profile_image"`
}

type AuthorAPIResponse struct {
	AuthorList []Author `json:"data"`
}
