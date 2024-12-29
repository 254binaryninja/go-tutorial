package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Goodmorning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("GoodBye and have a nice eveining %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, n := range n {
		f(n)
	}
}

// Function with return types
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	//cycleNames([]string{"Tony", "Tamara", "Warren", "Arnold", "Leon", "Peterson", "Brandon"}, sayGreeting)
	//cycleNames([]string{"Tony", "Tamara", "Warren", "Arnold", "Leon", "Peterson", "Brandon"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(12)

	fmt.Println("Circle Area is", a1, a2)
	fmt.Printf("Circle 1 area is %0.3f and circle 2 area is %0.3f\n", a1, a2)
}
