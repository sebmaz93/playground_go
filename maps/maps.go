package maps

import "fmt"

func main() {
	// you can have any type of KEY, and its dynamic compared to structs
	websites := map[string]string{
		"google": "google.com",
		"aws":    "amazon.com",
	}
	fmt.Println(websites)

	websites["fb"] = "facebook.com"
	fmt.Println(websites)

	delete(websites, "google")
	fmt.Println(websites)
}
