/*
 First, enter your wieght in kg👇
 Second, enter your height in m☝
 Then, you will get the result🤪
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

const mainTitle = "BMI Calculator"
const seprator = "---------------"
const weightPromptText = "Please enter your weight(kg): "
const heightPromptText = "Please enter your height(m): "

func main() {
	fmt.Println(mainTitle)
	fmt.Println(seprator)
	fmt.Print(weightPromptText)
	inputWeight, _ := reader.ReadString('\n')
    fmt.Printf("%v",inputWeight)
	fmt.Print(heightPromptText)
	inputHeight, _ := reader.ReadString('\n')
    fmt.Printf("%v",inputHeight)
	inputWeight = strings.Replace(inputWeight, "\n", "", -1)
	inputHeight = strings.Replace(inputHeight, "\n", "", -1)
	weight, _ := strconv.ParseFloat(inputWeight, 64)
	height, _ := strconv.ParseFloat(inputHeight, 64)
	bmi := weight / (height * height)
    fmt.Println("")
	fmt.Printf("Your BMI: %.2f", bmi)
}
