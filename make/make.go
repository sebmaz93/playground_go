package make

import "fmt"

// this is type alias
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)
	userNames = append(userNames, "seb")
	userNames = append(userNames, "jax")
	fmt.Println(userNames)

	ratings := make(floatMap, 3)

	ratings.output()
}
