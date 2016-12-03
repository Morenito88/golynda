package main

import (
	"fmt"
	"sort"
)

func main() {

	states := make(map[string]string)
	fmt.Println(states)

	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "OR") // Remove element
	fmt.Println(states)

	states["NY"] = "New York"

	for k, v := range states { // Loop map
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states { // Loop map
		keys[i] = k
		i++
	}

	sort.Strings(keys) // Sort map values
	fmt.Println("\nSorted")

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
