//A very simple calculator written in Go. 


package main

import "fmt"

func sum (b1, b2 float64) float64 {
  return (b2+b1)
}

func mult (c1, c2 float64) float64 {
  return (c1*c2)
}

func del (d1, d2 float64) float64 {
  return (d1/d2)
}

func min (e1, e2 float64) float64 {
  return (e1-e2)
}

func step (f1, f2 float64) float64 {
  var x float64 = 1
  var y = f1
  for x < f2 {
    f1 *= y
    x++
  }
  return (f1)
}

func main() {
  fmt.Println ("Enter the expression this way: \n x \n «operation» (+, -, *, /, ^) \n y")
  var a1 float64
  fmt.Scanln (&a1)
  var deystvie string
  fmt.Scanln (&deystvie)
  var a2 float64
  fmt.Scanln (&a2)
  var resu float64
  
  switch deystvie {
    case "+":
      resu = sum (a1, a2)
    case "*":
      resu = mult (a1, a2)
    case "/":
      resu = del (a1, a2)
    case "-":
      resu = min (a1, a2)
    case "^":
      resu = step (a1, a2)
  }
  
  fmt.Println (resu)
}
