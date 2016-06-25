## ButterCMS Go API Wrapper
This wrapper is meant to enable Go developers to quickly and easily get up and running with [ButterCMS](https://buttercms.com/). It is based of off the [API documentation](https://buttercms.com/docs/api/).

## Available API Methods
```
// Feeds
ButterCMS.GetFeed("feedType")

// Posts
ButterCMS.SearchPosts("query", page) // page should be an integer
ButterCMS.GetPosts()
ButterCMS.GetPost("slug")

// Authors
ButterCMS.GetAuthors(true|false) // To include recent posts or not
ButterCMS.GetAuthor("slug")

// Categories
ButterCMS.GetCategories(true|false) // To include recent posts or not
ButterCMS.GetCategory("slug")
```

## Usage
To include this package in your Go code, simply use `import "ButterCMS"` at the top of your file

## Example Usage
```
package main

import (
	"ButterCMS"
	"fmt"
)

func main() {
	ButterCMS.SetAuthToken("<auth_token>")

	resp, err := ButterCMS.GetAuthors()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", resp)
}

```
