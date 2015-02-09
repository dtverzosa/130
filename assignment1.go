package main
import "fmt"
type person struct {
name string
age int
}
func main() {
var s = person{"Dustin", 24}
fmt.Println(s.name)
fmt.Println(s.age)
var message string = "is"
b, c := 24, true
fmt.Println(s.name, message, b, c)

message2 := "This is the address of this string:"
var greeting *string = &message2
fmt.Println(message2, greeting)

}
