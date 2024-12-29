package main

import "fmt"

func main() {
	names := []string{"Tony", "Tamara", "Warren", "Arnold", "Leon", "Peterson", "Brandon"}

	for index, name := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}

		fmt.Printf("the value at pos %v is %v \n", index, name)
	}
}
