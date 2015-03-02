package main
import "fmt"
//Dustin Verzosa
//This program prints out the smallest number in a given array
func main() {

  x := []int{ 48,96,86,68, 57,82,63,70, 37,34,83,27, 19,97, 9,17, }
  y := x[0]
  i:= 1
  for i <= 15 {
    if y > x[i] {
      y = x[i]
    }
    i = i + 1
  }
  fmt.Println("Smallest number is", y)

}
