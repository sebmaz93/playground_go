package lists

import "fmt"

func main() {
	//var productNames = [5]string{"ball", "pen", "book", "chair", "desk"}
	//prices := [6]int{1, 2, 3, 4, 5, 6}
	//productNames[3] = "carpet"

	//fmt.Println(prices)
	//fmt.Println(productNames)

	//featPrices := prices[1:4]
	//fmt.Println(featPrices)
	//fmt.Println(len(featPrices), cap(featPrices)) // capacity is based on the array that we sliced from it because we can still chance our slice

	// dynamic array using slices
	//items := []int{10, 20}
	//fmt.Println(items)

	//updatedItems := append(items, 30, 40, 50) // append returns new slice and doesn't mutate
	//fmt.Println(updatedItems)

	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobbies := [3]string{"gaming", "reading", "drawing"}
	fmt.Println("Hobbies: ", hobbies)
	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(hobbies[0], "--", hobbies[1:])
	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	hobby1 := hobbies[:2]
	hobby2 := hobbies[0:2]
	fmt.Println("hobby1 hobby2", hobby1, hobby2)
	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	hobby1 = hobbies[1:]
	fmt.Println("hobby1 after", hobby1)
	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"land job", "master go", "develop"}
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "master golang"
	goals = append(goals, "improve")
	fmt.Println("goals", goals)
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	type product struct {
		title string
		id    int
		price int
	}

	products := []product{{
		"car",
		1,
		1000,
	}, {
		"moto",
		2,
		500,
	}}

	products = append(products, product{
		"bike",
		3,
		200,
	})

	fmt.Println(products)
}
