package main
import "fmt"
//Dustin Verzosa
/*This program prints out all the numbers from 1 to 100
Prints "Fizz" by multiples of 3, Prints "Buzz" by mutiples of 5
Prints "FizzBuzz" by multiples of both 3 and 5
*/
func main() {

  i:= 1
  for i <= 100 {
    if i % 3 == 0 && i % 5 == 0{
      fmt.Println(i, "FizzBuzz")
    } else if i % 3 == 0 {
      fmt.Println(i, "Fizz")
    } else if i % 5 == 0 {
      fmt.Println(i, "Buzz")
    } else {
      fmt.Println(i)
    }
    i = i + 1
  }

}
