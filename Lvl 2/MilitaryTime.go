package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "time"
    )

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanLines)
    
    fmt.Print("Input time (h:mm AM|PM): ")
    if scanner.Scan() {
        input := scanner.Text()
        
        mt, err := MilitaryTime(input)
        if err != nil {
            fmt.Fprint(os.Stderr, err)
            os.Exit(1)
        }
        fmt.Printf("24h time: %s", mt)
        return
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
}


func MilitaryTime(st string) (mt string, err error) {
    var dt time.Time
    // case insensitive support
    st = strings.ToUpper(st)
    // from 12h format
    dt, err = time.Parse("3:04 PM", st)
    if err == nil {
        // to 24h format
        mt = dt.Format("15:04")
    }
    return
}
