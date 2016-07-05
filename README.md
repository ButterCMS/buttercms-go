## ButterCMS Go API Wrapper
This wrapper is meant to enable Go developers to quickly and easily get up and running with [ButterCMS](https://buttercms.com/). It is based of off the [API documentation](https://buttercms.com/docs/api/).

## Available API Methods
```
// Feeds
ButterCMS.GetFeed("rss|atom|sitemap")

// Posts
params := map[string]string{"page": "2"}
ButterCMS.SearchPosts("query", params)
ButterCMS.GetPosts(params)
ButterCMS.GetPost("slug")

// Authors
params = map[string]string{"include": "recent_posts"}
ButterCMS.GetAuthors(params)
ButterCMS.GetAuthor("slug", params)

// Categories
ButterCMS.GetCategories(params)
ButterCMS.GetCategory("slug", params)

// Content Fields
ButterCMS.GetContentFields([]string{"content-field-slug"})
```

## Usage
To include this package in your Go code, simply use `import "ButterCMS"` at the top of your file

## Example Usage
```go
package main

import (
	"ButterCMS"
	"fmt"
)

func main() {
	ButterCMS.SetAuthToken("<auth_token>")

	params := map[string]string{"include": "recent_posts"}
	authorResp, err := ButterCMS.GetAuthors(params)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", authorResp)

	params = map[string]string{"author_slug": "author-slug"}
	postResp, err := ButterCMS.GetPosts(params)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", postResp)
}
```
