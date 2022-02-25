// You can access the values of the slice using the same syntax as with arrays:

package main

import "fmt"

func main() {
  a := [5]int{0, 2, 4, 6, 8}
  var s []int = a[1:3]  // a portion

  fmt.Println(s[1])   // a item by index
  
  // Note the same address of data
  fmt.Printf("s[1] %v  %p\n", s[1], &s[1])
  fmt.Printf("a[2] %v  %p\n", a[2], &a[2])
}
