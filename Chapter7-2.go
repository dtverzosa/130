ckage main
import "fmt"
//Dustin Verzosa
//This program calls func "half", which cuts an integer in have and returns the answer and 
//a boolean answer on whether or not the nmber was even
func half(x int) (int, bool) {
  y := false
  if x % 2 == 0 {
    y = true
  }
  return x/2, y
}

func main() {

  x := 17
  fmt.Println(x)
  fmt.Println(half(x))

}
