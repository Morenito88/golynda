package main

import (
	"fmt"
)

func main() {

	doSomething()

	sum := addValues(2, 53)
	fmt.Println("Sum: ", sum)

	sum = addAllValues(23, 21, 6)
	fmt.Println(sum)
}

func doSomething() { // Lowercase == private function
	fmt.Println("Doing somthing.")
}

func addValues(val1, val2 int) int {
	return val1 + val2
}

func addAllValues(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	fmt.Printf("%T\n", values)
	return sum
}
