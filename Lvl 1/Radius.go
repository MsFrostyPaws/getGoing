package main                                                               	 
import (                                                                   	 
	"fmt"                                                                  	 
		"math"                                                                 	 
		)                                                                          	 
func main() {                                                              	 
	var radius int                                                         	 
		var area float32                                                       	 
	fmt.Println("Finding the area of a circle")                                    	 
		fmt.Println("--------------------")                                    	 
			fmt.Print("Enter radius: ")                                            	 
				fmt.Scanf("%d", &radius)                                               	 
	area = math.Pi * (float32(radius) * float32(radius))                   	 
		fmt.Printf("Circle Area: %.2f\n", area)                                	 
		}
