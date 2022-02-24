// Sometimes it’s very useful to set reminder to help you accomplish all of your tasks.

package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanLines)
    
    fmt.Print("Enter event name: ")
    var event string
    if scanner.Scan() {
        event = scanner.Text()
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
    fmt.Println()

    setReminder(event)
}

func DefaultIfBlank(s string, fallback string) string {
    if IsBlank(s) {
        s = fallback
    }
    return s
}

func IsBlank(s string) bool {
    s = strings.TrimSpace(s)
    return len(s) == 0
}

func setReminder(event string) {
    event = DefaultIfBlank(event,
        "!Unknown Event!")
    fmt.Println("You set a reminder about", event)
}
