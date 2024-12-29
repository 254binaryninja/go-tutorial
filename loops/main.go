package main

import (
	"fmt"
	"sort"
)

func main() {
	//x := 0
	//
	//for x < 5 {
	//	fmt.Println("The value of X is :", x)
	//	x++
	//}

	//for i := 0; i < 10; i++ {
	//	fmt.Println("The value of i is :", i)
	//}

	names := []string{"Tony", "Tamara", "Warren", "Arnold", "Leon", "Peterson", "Brandon"}
	sort.Strings(names)

	//for i := 0; i < len(names); i++ {
	//	fmt.Println("The value at the current index is :", names[i])
	//}

	// For a value you do not want to use replace with an _(underscore)
	for index, value := range names {
		fmt.Printf("The index at the current index is :%v while the value is :%v \n", index, value)
	}
}
