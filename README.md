# in-memory cache
=================

In-memory cache with the following methods:

*Set(key string, value interface{})*

*Get(key string)*

*Delete(key)*

See it in action:
## Example

```go
package main

import (
	"fmt"

	"github.com/vsakivskiy/cache"
)

func main() {
	c := cache.NewCache()
	c.Set("uuid", "4dd76b1c-b7a7-11ed-afa1-0242ac120002")
	c.Set("userId", 88)

	uuid, _ := c.Get("uuid")
	fmt.Println(uuid)

	userId, _ := c.Get("userId")
	fmt.Println(userId)

	c.Delete("userId")
	userId, err := c.Get("userId")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(userId)
	}
}
```
