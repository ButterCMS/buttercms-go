## ButterCMS Go API Wrapper
This wrapper is meant to enable Go developers to quickly and easily get up and running with [ButterCMS](https://buttercms.com/).

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
