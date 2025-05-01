package fundamentals

import (
	"fmt"
	foo "net/http"
)

func main() {
	fmt.Println("Go Standard Library")

	res, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("HTTP Response Status:", res.Status)
}
