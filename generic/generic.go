package generic

func add[T int | float64](a, b T) T {
	return a + b
}

func main() {
	add[float64](1, 2)
}
