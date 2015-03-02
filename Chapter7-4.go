package main
import "fmt"
//Dustin Verzosa
//Fibonacci Sequence
func fib(x uint) uint {
    if x == 0 {
        return 1
    }
    return x + fib(x-1)
}
func main() {
    fmt.Println(fib(5))
}
