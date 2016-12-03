package main

import (
	"fmt"
)

func main() {

	defer fmt.Println("Close the file!") // Defer the statement to the last position of code execution
	fmt.Println("Open the file!")

	defer fmt.Println("St1")
	defer fmt.Println("St2")

	MyFunc()
	defer fmt.Println("St3")
	defer fmt.Println("St4")
	fmt.Println("St undefer")
}

func MyFunc() {

	defer fmt.Println("Defer in func!")
	fmt.Println("Not defer in func!")
}
