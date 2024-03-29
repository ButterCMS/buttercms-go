# ButterCMS Go API Wrapper

This wrapper is meant to enable Go developers to quickly and easily get up and running with [ButterCMS](https://buttercms.com/). It is based of off the [API documentation](https://buttercms.com/docs/api/).

## Versioning

Each revision of the binding is tagged and the version is updated accordingly.

It is highly recommended you use `go modules` for vendoring/package management, in order to ensure a newer
version of the binding does not affect backwards compatibility.

## Installation

```sh
go get github.com/buttercms/buttercms-go
```

## Pages

Full list of `params` availble in our official [API Documentation](https://buttercms.com/docs/api/?go#pages)

```
params := map[string]string{"foo": "bar"}
ButterCMS.GetPages("news", params)
ButterCMS.GetPage("news", "another-test-page", params)
```

## Collections

Full list of `params` availble in our official [API Documentation](https://buttercms.com/docs/api/?go#collections)

```
params := map[string]string{"locale": "en"}
ButterCMS.GetContentFields([]string{"collection_key"}, params)
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

### Preview mode

Preview mode can be used to setup a staging website for previewing content fields or for testing content during localdevelopment. To fetch content from preview mode call `ButterCMS.SetPreviewMode` with `true` input:

```go
package main
...
ButterCMS.SetPreviewMode(true)
```

## Example Usage

```go
package main

import (
	"github.com/buttercms/buttercms-go"
	"fmt"
	"encoding/json"
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

	contentParams := map[string]string{}
	contentKeys := []string{"somethings"}
	contentResp, err := ButterCMS.GetContentFields(contentKeys, contentParams)
	if err != nil {
		panic(err.Error())
	}

	j, err := json.Marshal(contentResp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j))
}
```

### Other

View these Go [Blog engine](https://buttercms.com/golang-blog-engine/) and [Full CMS](https://buttercms.com/golang-cms/) for other examples of using ButterCMS with Go.
