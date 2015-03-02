package main
import "fmt"
//Dustin Verzosa
//This program prints out all the multiples of 3 from 1 to 100
func main() {

  i:= 1
  for i <= 100 {
    if i % 3 == 0 {
      fmt.Println(i)
    }
    i = i + 1
  }

}
