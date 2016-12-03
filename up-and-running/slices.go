package main

import (
	"fmt"
	"sort"
)

func main() {

	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "Black") // Add new element
	fmt.Println(colors)

	colors = append(colors[1:]) // Remove first element
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1]) // Remove last element
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 65
	numbers[1] = 34
	numbers[2] = 32
	numbers[3] = 52
	numbers[4] = 123
	fmt.Println(numbers)

	numbers = append(numbers, 231)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))

	sort.Ints(numbers)
	fmt.Println(numbers)
}
