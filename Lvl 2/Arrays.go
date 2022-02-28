package main
import "fmt"

//get average
func average(nums ...int) int{
    sum := 0
    var size int = 0
    
    for i,n := range nums{
        sum += n
        size = i + 1
        
    }
    
    result := sum/size
    return result
}

//makes array in given range
func rangearr(first int,last int,skip int) []int{
    if skip == 0{
        skip += 1
    }
    if last == 0{
        last += 10 + first
    }
    arr := []int{}
    if last > first{
        for s:=first;s<=last;s+=skip{
            arr = append(arr,s)
        }
    }
    return arr
}

//gets total sum of array
func total(nums ...int) int{
    sum := 0
    for _,v := range nums{
        sum += v
    }
    return sum
}

//gets the size of array
func size(nums ...int) int{
    size := 0
    for i,_ := range nums{
        size = i + 1
    }
    return size
}

//get max
func max(nums ...int) int{
    max := 0
    for _,v := range nums{
        if v > max{
            max = v
        }
    }
    return max
}

//get min
func min(nums ...int) int{
    execution := 0
    min := 0
    for _,v := range nums{
        if execution == 0{
            min = v
            execution++
        }else{
            if min > v{
                min = v
            }
        }
    }
    return min
}

//execute
func main() {
    arr := []int{56,89,34,123,45,68,95}
    average := average(arr...)
    max := max(arr...)
    min := min(arr...)
    total := total(arr...)
    size := size(arr...)
    ranged := rangearr(min,max,1)
    
    fmt.Println("All you need to do is just edit the arr variable")
    fmt.Println("")
    fmt.Println("array")
    fmt.Println(arr)
    fmt.Println("")
    fmt.Println("average of array")
    fmt.Println(average)
    fmt.Println("")
    
    fmt.Println("largest number in the array")
    fmt.Println(max)
    fmt.Println("")
    
    fmt.Println("smallest number in array")
    fmt.Println(min)
    fmt.Println("")
    
    fmt.Println("total sum of array")
    fmt.Println(total)
    fmt.Println("")
    
    fmt.Println("total size of array")
    fmt.Println(size)
    fmt.Println("")
    
    fmt.Println("array made by ranging   smallest number in array to largest number in array")
    fmt.Println(ranged)
    fmt.Println("")
}
