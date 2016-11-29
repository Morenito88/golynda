package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum:", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum:", floatSum)

	var v1, v2, v3, bigSum big.Float
	v1.SetFloat64(23.5)
	v2.SetFloat64(65.1)
	v3.SetFloat64(76.3)

	bigSum.Add(&v1, &v2).Add(&bigSum, &v3)
	fmt.Printf("Big Sum = %.10g\n", &bigSum)

	circleRadius := 15.5
	circumference := circleRadius * math.Pi
	fmt.Printf("circumference: %.2f\n", circumference)
}
