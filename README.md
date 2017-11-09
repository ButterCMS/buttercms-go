# ButterCMS Go API Wrapper

This wrapper is meant to enable Go developers to quickly and easily get up and running with [ButterCMS](https://buttercms.com/). It is based of off the [API documentation](https://buttercms.com/docs/api/).

## Versioning

Each revision of the binding is tagged and the version is updated accordingly.

Given Go's lack of built-in versioning, it is highly recommended you use a
[package management tool](https://github.com/golang/go/wiki/PackageManagementTools) in order to ensure a newer
version of the binding does not affect backwards compatibility.

## Installation

```sh
go get github.com/buttercms/buttercms-go
```

## Pages
```
params := map[string]string{"foo": "bar"}
ButterCMS.GetPages("news", params)
ButterCMS.GetPage("news", "another-test-page", params)
```

## Content Fields
```
ButterCMS.GetContentFields([]string{"content-field-slug"})

// Test Mode
params := map[string]string{"test": 1}
ButterCMS.GetContentFields([]string{"content-field-slug"}, params)

// Localization
params := map[string]string{"locale": "en"}
ButterCMS.GetContentFields([]string{"content-field-slug"}, params)
```

## Blog Engine
```
// Posts
params := map[string]string{"page": "2"}
ButterCMS.GetPosts(params)
ButterCMS.GetPost("slug")

// Authors
params = map[string]string{"include": "recent_posts"}
ButterCMS.GetAuthors(params)
ButterCMS.GetAuthor("slug", params)

// Categories
ButterCMS.GetCategories(params)
ButterCMS.GetCategory("slug", params)

// Feeds
ButterCMS.GetFeed("rss|atom|sitemap")

// Search
ButterCMS.SearchPosts("query", params)
```

## Usage

To include this package in your Go code, simply use `import "ButterCMS"` at the top of your file

## Example Usage

```go
package main

import (
	"github.com/buttercms/buttercms-go"
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

	pagesResp, err := ButterCMS.GetPages("news", nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%d Results\n", pagesResp.MetaData.Count)
	for index, page := range pagesResp.PageList {
		fmt.Printf("Result %d: %s\n", index, page.Slug)

		if (len(page.Fields) > 0) {
			fmt.Println("Custom fields:")
			for fieldName, fieldValue := range page.Fields {
				switch value := fieldValue.(type) {
					case string:
						fmt.Printf("Field '%1s' has value '%2s'\n", fieldName, value)
					default:
						fmt.Println("Other type\n")
				}
			}
		}
	}
}
```

### Other

View these Go [Blog engine](https://buttercms.com/golang-blog-engine/) and [Full CMS](https://buttercms.com/golang-cms/) for other examples of using ButterCMS with Go.
