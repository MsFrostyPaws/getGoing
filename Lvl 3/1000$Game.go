package main

import (
         "fmt"
         "bufio"
         "log"
         "os"
         "strconv"
         "strings"
         "math/rand"
         "time"
)

func main() {
       seconds := time.Now().Unix()
       rand.Seed(seconds)
       target := rand.Intn(100) + 1
       fmt.Println("try your luck!!! you'll win 1000$ if you guess the right number between 1 and 100.")
       fmt.Println("wanna try ??")

       reader := bufio.NewReader(os.Stdin)

       success := false
       for guesses := 0; guesses < 10; guesses++ {
              fmt.Println("You have", 10-guesses, "guesses left.")

              fmt.Print("Make a Guess: ")
              input, err := reader.ReadString('\n')
              if err != nil {
                     log.Fatal(err)
               }
              input = strings.TrimSpace(input)
              guess, err := strconv.Atoi(input)
              if err != nil {
                     log.Fatal(err)
              }
              if guess < target {
                     fmt.Println("your guess was LOW.")
              } else if guess > target {
                     fmt.Println("your guess was HIGH.")
              } else {
                    success = true
                    fmt.Println("wow you won congratulations $$$")
                    break
              }
      }
   
      if !success {
             fmt.Println("sorry, you didn't guess my number. it was:", target)
      }
}
