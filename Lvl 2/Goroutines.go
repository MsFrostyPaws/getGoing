/*
Goroutines

When running concurrent programs, do not often want to wait for one task to finish before starting a new one.
To achieve concurrency, let's start the function calls as Goroutines, using the go keyword:

*/

package main
import (
    "fmt"
    "log"
    "os"
    "time"
)

/*
The out() function simply outputs numbers in the given range. We use a time.Sleep() to emulate work being done between the outputs just for demonstration purposes. It simply waits for the provided time (50ms) before continuing the execution.
*/

func out(from, to int) {
    for i:=from; i<=to; i++ {
        time.Sleep(50 * time.Millisecond)
        fmt.Println(from, to, i)
        log.Println(from, to, i)
    }
}

/*
As you can see, it is very simple to start a Goroutine -- we simply need to add a go keyword before the function call.

If we run the program, we get a surprising output -- it will result in No Output.
But why? This output happens because the main() function exits before the Goroutines complete.
*/

func main() {
    go out(0, 5)
    go out(6, 10)
}
func init() {
    log.SetFlags(log.Ltime|log.Lmicroseconds|log.LUTC)
    log.SetOutput(os.Stderr)
}

/*
Our program has 3 virtual threads: the 2 function calls, and main(). Our 2 function calls get executed concurrently, and main() does not wait for them to finish.
*/
