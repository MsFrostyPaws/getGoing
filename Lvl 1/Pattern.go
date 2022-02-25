package main

import (
    ."fmt"
)

func main() {
    pattern := [5]int{1,2,3,4,5}
    for i , _ := range pattern {
        Println(pattern[:i+1])
    }
}
