package main

import "fmt"

/* Insert a number (Celsius) to convert to Fahrenheit*/

func main(){
    var input float64
    fmt.Scanln(&input)
    
    fmt.Println(input,"°C =",toFahrenheit(input), "°F")
}

func toFahrenheit(celsius float64) float64 {
    return ((celsius * 9/5) + 32)
}
