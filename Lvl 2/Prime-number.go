package main

import "fmt"

func main() {
   var inp int;
  var i,cout int = 1,0;
  fmt.Scanln(&inp)
   
    for ( i<=inp){
if(inp%i==0){
cout++;
fmt.Println(inp,"/",i,"=",inp/i)
}
i++
}
if(cout==2){

 fmt.Println (inp,"is a prime number")
}
if(cout!=2){
    fmt.Println(inp,"is not a prime number")
    }

}
