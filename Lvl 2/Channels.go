/*
Channels
=====================

Let's use a channel to send data from a Goroutine and use it in main().

Our program needs to calculate and output the sum of all even numbers in a given range plus the sum of their squares and output the result:

     output = evenSum+squareSum

We will use two Goroutines: one to calculate the evenSum, and the other to calculate the squareSum.

We will get the data using channels in main(), then calculate and output the final sum.
*/

package main
import (
    "fmt"
    "strings"
)


func evenSum(from, to int, ch chan int) {
    ConditionalRangerEmmiter(ch, from, to, IsEven, Identity)
}
func squareSum(from, to int, ch chan int) {
    ConditionalRangerEmmiter(ch, from, to, IsEven, Square)
}

/*
As you can see, our functions send the result via channels.


Now we can call them as Goroutines in main() and output the resulting sum:
*/

func main() {
    var output int
    
    fmt.Print("Use a shared data channel? (y/n): ")
    var flag string
    if fmt.Scan(&flag); strings.ToLower(flag) == "y" {
        fmt.Println(`
USING SHARED CHANNEL....`)
        
        ch := make(chan int)
        defer close(ch)
        
        go evenSum(0, 100, ch)
        go squareSum(0, 100, ch)
        
        // wait 2 times for channel data
        output = <-ch + <-ch
        
    } else {
        fmt.Println(`
USING SEPARATED CHANNELS....`)

        var evenCh, sqCh chan int
        evenCh, sqCh = make(chan int), make(chan int)
        defer close(evenCh)
        defer close(sqCh)
        
        go evenSum(0, 100, evenCh)
        go squareSum(0, 100, sqCh)
        
        // wait for data of both channels
        output = <-evenCh + <-sqCh
    }
    
    fmt.Println(`
    Result:`, output)
}

/*
If you do not need to send data to a channel anymore, you can close it using close(ch), where ch is the name of the channel. This is done in the sender.
*/




//
// Async Iterators
//

func ConditionalRangerEmmiter(
        receiver   chan int,
        from, to   int,
        predicate  IntPredicate,
        operation  IntOperator,
        ) {
    receiver <- ConditionalRanger(
            from, to, predicate, operation)
}

//
// Iterators
//

func ConditionalRanger(
        from, to   int,
        predicate  IntPredicate,
        operation  IntOperator,
        ) (result int) {
    for n := from; n <= to; n++ {
        if predicate(n) {
            result += operation(n)
        }
    }
    return
}


//
// IntFunctors
//

type IntOperator  func(int) int

func Identity(value int) int {
    return value
}
func Square(value int) int {
    return value*value
}


//
// IntPredicates
//

type IntPredicate func(int) bool

func IsEven(value int) bool {
    return value & 1 == 0 // fast remainder moving bits
//    return value % 2 == 0
}
func Always(_ int) bool {
    return true
}
func Never(_ int) bool {
    return false
}
