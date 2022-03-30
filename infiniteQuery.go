package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//Insert your code here
	var num string
	for {
		fmt.Println("Please enter a  number: ")
		fmt.Scanln(&num)

		if number, _ := strconv.ParseInt(num, 10, 0); number%2 == 0 {
			fmt.Printf("%d is an even number ", number)
		} else {
			fmt.Printf("%d is a odd number ", number)
		}

		if number, _ := strconv.ParseInt(num, 10, 0); math.Abs(float64(number)) > 9 {
			fmt.Println("and has more than 1 digit.")
		} else {
			fmt.Println("and has 1 digit.")
		}
	}
}
