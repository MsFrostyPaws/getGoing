package main
import "fmt"

type Timer struct {
    id string
    value int 
}

func (ptr *Timer) tick(val int) {
    ptr.value = val
}

func main() {
    var x int 
    fmt.Scanln(&x)

    t := Timer{"timer1", 0}
    
    for i:=0;i<x;i++ {
        t.tick(i+1)
        fmt.Println(t.value)  
    }
}
