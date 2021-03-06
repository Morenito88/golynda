package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	/*dow := rand.Intn(6) + 1
	fmt.Println("Day", dow)*/

	result := ""

	switch dow := rand.Intn(6) + 1; dow {
	case 1:
		result = "Its Sunday"
	case 7:
		result = "Its Saturday"
	default:
		result = "Its a weekday"
	}

	fmt.Println(result)

	x := 42
	switch {
	case x < 0:
		result = "Less than zero"
		//fallthrough //continue to the next case (dont break it)
	case x == 0:
		result = "Equal zero"
	default:
		result = "Greater than zero"

	}

	fmt.Println(result)
}
