package main

import (
    "fmt"
    "bufio"
    "os"
    "encoding/json"
)

func main(){

    var name string
    var addr string

    person := make(map[string]string)

    read_input := bufio.NewScanner(os.Stdin)

    fmt.Printf("\nPlease enter a name: ")
    read_input.Scan()
    name = read_input.Text()

    fmt.Printf("Please enter an address: ")
    read_input.Scan()
    addr = read_input.Text()

    person["name"] = name
    person["address"] = addr

    jsonperson, _ := json.Marshal(person)
    fmt.Println("Your JSON Object is: ")
    fmt.Println(string(jsonperson))
}
