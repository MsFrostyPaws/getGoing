package main

import "fmt"

func pow(a,b int)int{
    var c int = a
    if b==0{
        return 1
    }else{
    for i:=0;i<b-1;i++{
        a*=c
      }
      return a
    }
}

func isPalindrome(num int)bool{
    var num1,rem,digit,i int = 0,num,0,0
    for true{
        rem/=10
        i++
        if rem==0{
            rem=num
            break
        }
    }
    
    for j:=i-1;j>=0;j--{
        digit=rem%10
        num1+=digit*(pow(10,j))
        rem/=10
    }

    if num==num1{
        return true
    }else{
        return false
    }
}

func main(){
    fmt.Println("Hi friends!")
    var num int
    fmt.Println("Enter a number :")
    fmt.Scanln(&num)
    fmt.Println(num)
    if isPalindrome(num){
        fmt.Println(" Given number is a Palindrome.")
    }else{
        fmt.Println(" Given number is not a Palindrome.")
    }
}
