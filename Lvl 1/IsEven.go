

package main

import "fmt"

func main() {

  fmt.Print("Enter number: ")
  
  var input int
  
  fmt.Scanln(&input)
  
  fmt.Print(input)
  
  if(input % 2 == 0){
	  
	  fmt.Print(" is even\n")
	  
  }else{
	  
	  fmt.Print(" is odd\n")
  }
  
}
