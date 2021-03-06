// check someone’s age...not just at the bar!

package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    var text string
    fmt.Print("Input your age: ")
    if _, err := fmt.Scanln(&text); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    
    age, err := Atoui(text)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    
    if isAdult(age) {
        fmt.Println("You are considered an adult")
    } else {
        fmt.Println("Too pretty young to be adult.")
    }
}

func isAdult(age uint) bool {
    return age >= 18
}

func Atoui(s string) (uint, error) {
    ui64, err := strconv.ParseUint(s, 10, 0)
    return uint(ui64), err
}
