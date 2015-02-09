package main
import "fmt"
//Dustin Verzosa
//constant PI is pi without the decimal places to keep it an int
const (
PI = 3
)

//calculates diameter and general area
func calculate(radius, PI int) (int, int) {
return radius * 2, radius * PI * radius
}

func main() {

circleradius := 2
fmt.Println(PI)
fmt.Println(calculate(circleradius, PI))


}
