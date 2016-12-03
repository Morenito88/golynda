package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "An implicit string"
	fmt.Printf("str1: %v:%T\n", str1, str1)

	str2 := "An explicit string"
	fmt.Printf("str2: %v:%T\n", str2, str2)

	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str1))

	lValue := "hello"
	uValue := "HELLO"
	fmt.Printf("lValue = uValue? %v\n", (lValue == uValue))
	fmt.Printf("lValue = uValue? %v\n", strings.EqualFold(lValue, uValue))

	fmt.Println("Contains? ", strings.Contains(str1, "exp"))
	fmt.Println("Contains? ", strings.Contains(str2, "exp"))
}
