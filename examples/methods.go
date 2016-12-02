package main

import (
	"fmt"
)

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speak() { // Method of Dog struct
	fmt.Println(d.Sound)
}

func (d *Dog) SpeekThreeTimes() { // Pointer to field Sound
	d.Sound = fmt.Sprintf("%v! %v! %v!", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

func main() {

	poodle := Dog{"Poodle", 37, "Woof"}
	fmt.Println(poodle)
	poodle.Speak()

	poodle.Sound = "Arf"
	poodle.Speak()

	poodle.SpeekThreeTimes()
	poodle.SpeekThreeTimes()
}
