package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	numbers := []int{}
	for {
		fmt.Println("Enter The float number : ")
		var input string

		fmt.Scanln(&input)
		if strings.ToLower(input) == "x" {
			fmt.Println("End !!!")
			break
		}
		number, err := strconv.Atoi(input)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		numbers = append(numbers, number)
		sort.Ints(numbers)
		fmt.Println(numbers)
	}

}

// func main() {
// 	x := [...]int{4, 8, 5}
// 	fmt.Println(x)
// 	y := x[0:2]
// 	fmt.Println(y)
// 	z := x[1:3]
// 	fmt.Println(z)
// 	y[0] = 1
// 	z[1] = 3
// 	fmt.Print(x)
// }

// func main() {
// 	x := [...]int{1, 2, 3, 4, 5}
// 	y := x[0:2]
// 	z := x[1:4]
// 	fmt.Print(len(y), cap(y), len(z), cap(z))
// }

// func main() {
// 	x := map[string]int{
// 		"ian": 1, "harris": 2}
// 	for i, j := range x {
// 		if i == "harris" {
// 			fmt.Print(i, j)
// 		}
// 	}
// }

// type P struct {
// 	x string
// 	y int
// }

// func main() {
// 	b := P{"x", -1}
// 	a := [...]P{P{"a", 10},
// 		P{"b", 2},
// 		P{"c", 3}}
// 	for _, z := range a {
// 		if z.y > b.y {
// 			b = z
// 		}
// 	}
// 	fmt.Println(b.x)
// }

func main() {
	s := make([]int, 0, 3)
	s = append(s, 100)
	fmt.Println(len(s), cap(s))
}
