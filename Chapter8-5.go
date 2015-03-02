package main
import "fmt"
//Dustin Verzosa
//Swap using pointers

func swap(x *int, y *int) {
  var z int
  z = *x
  *x = *y
  *y = z
}

func main() {

  x := 2
  y := 16
  swap(&x, &y)
  fmt.Print("x: ", x)
  fmt.Print(", y: ", y)

}
