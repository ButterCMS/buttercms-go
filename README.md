## ButterCMS Go API Wrapper
This wrapper is meant to enable Go developers to quickly and easily get up and running with [ButterCMS](https://buttercms.com/). It is based of off the [API documentation](https://buttercms.com/docs/api/).

## Available API Methods
```
// Feeds
ButterCMS.GetFeed("feedType")

// Posts
ButterCMS.GetPosts()
ButterCMS.SearchPosts("query")

// Authors
ButterCMS.GetAuthors()

// Categories
ButterCMS.GetCategories()
```

## Usage
To include this package in your Go code, simple use `import "ButterCMS"` at the top of your file

## Example Usage
```
package main

import (
	"ButterCMS"
	"fmt"
)

func main() {
	ButterCMS.SetAuthToken("<auth_token>")

	authorResp, err := ButterCMS.GetAuthors()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", authorResp)
}

```
