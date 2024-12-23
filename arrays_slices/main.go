package main

import "fmt"

func main() {
	//Arrays use curly braces and also typed on the right sides
	//var arr [5]int = [5]int{1, 2, 3, 4, 5}

	//shorthand
	var arr = [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr, len(arr))

	// Slices (uses arrays under the hood)
	var scores = []int{1, 2, 3, 4, 5}
	// This method also applies to arrays in go
	scores[2] = 100

	scores = append(scores, 1000)
	fmt.Println(scores, len(scores))
}
