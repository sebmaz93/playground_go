package loops

import "fmt"

func Main() {
	userNames := []string{"seb", "jax", "tom"}

	for idx, user := range userNames {
		fmt.Println("index: ", idx)
		fmt.Println("username: ", user)
	}
}
