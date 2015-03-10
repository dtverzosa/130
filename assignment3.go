package main

import "fmt"

func myConditional(myCheck bool) {

  myFavName := "Dustin"
  if  myCheck {
    fmt.Println(myFavName + " is my name")
  } else {
    fmt.Println(myFavName + " is not my name")
  }

}
///////////Main function Starts here//////////////
func main() {

  myname := "Dustin"
  var dtv bool
  switch myname {
    case "Dustin":
      fmt.Println("Verzosa")
    case "Bryce":
      fmt.Println("Johnson")
    case "Boogie":
      fmt.Println("Romero")
  }

  if myname == "Dustin"{
    dtv = true
  } else {
    dtv = false
  }
  
  myConditional(dtv)

}
