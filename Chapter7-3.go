package main
import "fmt"
//Dustin Verzosa
//This program uses a variadic parameter to determine the largest number in a given list of numbers
func add(args ...int) int {
    biggest := 0
    for _, v := range args {
        if v > biggest {
          biggest = v
        }
    }
    return biggest
}
func main() {
    fmt.Println(add(1,2,3,100,2))
}
