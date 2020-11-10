package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a string  : ")
	var input string
	if scanner.Scan() {
		input = scanner.Text()
	}
	fmt.Println(input)

	if strings.HasPrefix(strings.ToLower(input), "i") && strings.Contains(strings.ToLower(input), "a") && strings.HasSuffix(strings.ToLower(input), "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
