package main

import "fmt"

func main() {
	// i, _ := strconv.Atoi("10")
	// y := i * 2
	// fmt.Println(y)
	var y *int
	var x int
	z := 3

	y = &z
	x = *y

	fmt.Println(x)

	fmt.Println(y)
}
