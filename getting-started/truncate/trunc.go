package main 
  
import "fmt"

func main() {
   fmt.Println("Enter The float number : ")
   var number float32
   fmt.Scanln(&number)
   fmt.Print("Truncate number : ")
   fmt.Println(int(number)) 

}
