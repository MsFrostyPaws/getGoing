package main

import (
    "fmt"
    "math/rand"
    "time"
)

func bubbleSort(arr[]int) []int {
    swapped := true

    for swapped {
        swapped = false
        for i := 0; i < len(arr) - 1; i++ {
            if arr[i] > arr[i + 1] {
                arr[i], arr[i + 1] = arr[i + 1], arr[i]
                swapped = true
            }
        }
    }

    return arr
}

func main() {
    rand.Seed(time.Now().UnixNano()) // Seed for rand
    randArr := rand.Perm(20)         // Perm function from rand, generate random n-size array
    
    // Print results
    fmt.Println("Before sorting: ")
    fmt.Println(randArr)
    fmt.Println("After sorting: ")
    fmt.Println(bubbleSort(randArr))
}
