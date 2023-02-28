# in-memory-cache
================================

In-memory cache with the following methods:

**_ Set(key string, value interface{}) _**
**_ Get(key string) _**
**_ Delete(key) _**

See it in action:
## Example

```go
func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId := cache.Get("userId")

	fmt.Println(userId)
}
```
