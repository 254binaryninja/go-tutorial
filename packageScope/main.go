package main

import "fmt"

// Always run this together with the file it depends on
func main() {
	sayHello("Arnold")

	for _, v := range points {
		fmt.Println(v)
	}
}
