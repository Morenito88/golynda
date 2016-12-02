package main

import (
	"fmt"
)

func main() {

	//var x float64 = 42
	var result string

	if x := 42; x < 0 {
		result = "Less than zero"
	} else if x == 0 {
		result = "Equal zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println(result)

}
