package main

import "fmt"

type Person struct {
  name string
  age string 
  height string 
  weight string 
}
type Car struct {
    brand string
    model string
    color string
}
type Residence struct {
    house string
    lndmark string
    city string
    country string
}

func main() {
  x := Person{"Samuel","26", "168.2", "63"}
  y := Car{"Lexus", "X350", "black"}
  z := Residence{"Duplex", "Victoria Island", "Lagos", "Nigeria"}
  
//loop printing 
 var bio string 
 bio = "Mr." 
 bio += x.name 
 bio += " is " 
 bio += x.age
 bio += " years old, he is " 
 bio += x.height 
 bio += "cm tall and weigh " 
 bio += x.weight 
 bio += "kg."
  fmt.Println(bio)
  fmt.Println("\n")
  
  fmt.Println("He has a " + y.brand + y.model + " car colored " + y.color + ".")
  
  fmt.Println("\n")

  fmt.Println("He lives in a " + z.house + " at " + z.lndmark + " in " + z.city + " city, " + z.country + ".")
}
