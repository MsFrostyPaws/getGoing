package main

import "fmt"

const Pi = 3.142

func main() {
    circle := Circle{Radius: 2.5}
    rect := Rectangle{Length: 31.2, Breath: 12.5}
    CalculateArea(circle)
    CalculateArea(rect)
}

type Shape interface {
    Area()
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() {
    area := Pi * c.Radius * c.Radius
    fmt.Printf("Area of circle is: %v\n", area)
}

type Rectangle struct {
    Length float64
    Breath float64
}

func (r Rectangle) Area() {
    area := r.Length * r.Breath
    fmt.Printf("Area of rectangle is: %v\n", area)
}


func CalculateArea(shape ...Shape) {
    for _, s := range shape {
        s.Area()
    }
}
