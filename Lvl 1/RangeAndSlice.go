

package main

import "fmt"

func main() {
  a := make([]int, 5)
  a[1] = 2
  a[2] = 3
  fmt.Println("slice", a, len(a), cap(a))

// During each iteration of the loop, it returns two values: the index of the element and its value.

  fmt.Println("range index-value")
  for i, v := range a {
    fmt.Println(i, v)
  }

// If you want only the values, you can skip the index using an underscore:

  fmt.Println("range omit index using underscore token")
  for _, v := range a {
    fmt.Println(v)
  }

  fmt.Println("range omit value")
  for i := range a {
    fmt.Println(i)
  }
  
// also can use indexed loop. also is more flexible to indicate when to stop and the start or increment.

  fmt.Println("for-i")
  for i, l := 0, len(a); i < l; i++ {
    fmt.Println(i, a[i])
  }
  
  fmt.Println("for-i odds")
  for i, l := 1, len(a); i < l; i+=2 {
    fmt.Println(i, a[i])
  }

  fmt.Println("for-i reverse")
  for i := len(a) - 1; i >= 0; i-- {
    fmt.Println(i, a[i])
  }
  
// it is safe iterate over a nil slice
  a = nil
  fmt.Println("range nil", a)
  for _, v := range a {
    fmt.Println(v)
  }
  fmt.Println("for-i nil", a, len(a))
  for i, l := 0, len(a); i < l; i++ {
    fmt.Println(i, a[i])
  }

}

// You can use ranges for slices as well as arrays.
// But in next chapters you learn that also for maps and channels.
