package main
import "fmt"

type Cart struct {
    prices []float32
}
func (x Cart) show() {
    var sum float32 = 0.0
    //calculate the sum of all prices in the Cart
    for _,p := range x.prices{
        sum += p
    }
    fmt.Println(sum)
}

func main() {
  c := Cart{}
  var n int
  fmt.Scanln(&n)
  
  // take n inputs and add them to the Cart
  var p float32
  var a int
  for a=0;a<n;a++{
     fmt.Scanln(&p)
      c.prices=append(c.prices,p)
  }
  //call the show() method of the Cart
  c.show()
}
