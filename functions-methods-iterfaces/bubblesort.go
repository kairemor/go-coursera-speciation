
package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func swap (slice []int, i int){

    one := slice[i+1]
    two := slice[i]

    slice[i] = one
    slice[i+1] = two
}

func BubbleSort (slice []int){

    sorted := false

    for sorted == false{

        swapped := false

        for i:=0; i < (len(slice)-1); i++{

            if slice[i] > slice[i+1]{
                swap(slice, i)
                swapped = true
            }
        }

        if swapped == false{
            sorted = true
        }
    }
}

func main(){

    read_input := bufio.NewScanner(os.Stdin)
    fmt.Println("\nPlease enter a sequence of up to ten integers separated by spaces.")
    fmt.Println("Hit enter when finished:")
    read_input.Scan()
    line := read_input.Text()

    user_ints := strings.Split(line, " ")

    var user_slice []int

    for i:=0; i < len(user_ints); i++ {
        user_int, err := strconv.Atoi(user_ints[i])
        if err == nil {
            user_slice = append(user_slice, user_int)
        } else {
            fmt.Println("Invalid input. Please run the program again with ")
            fmt.Println("a sequence of up to ten integers separated by spaces. ")
            fmt.Print("Exiting program ")
            os.Exit(1)
        }

    }

    BubbleSort(user_slice)

    fmt.Println("Sorted slice of integers:")

    for i:=0; i < len(user_slice); i++{
        fmt.Printf("%d ", user_slice[i])
    }

    fmt.Println()
}
