package main

import (
	"fmt"
	"sort"
)

func main() {
	//greeting := "hello their friends"
	//fmt.Println(strings.Contains(greeting, "hello"))
	//fmt.Println(strings.Index(greeting, "hello"))
	//fmt.Println(strings.LastIndex(greeting, "hello"))
	//fmt.Println(strings.Split(greeting, " "))

	////Sort package
	//ages := []int{45, 56, 78, 90, 47, 30, 4}
	//
	//sort.Ints(ages)
	//fmt.Println(ages)
	//
	////Finding position
	//index := sort.SearchInts(ages, 45)
	//fmt.Println(index)

	//Sorting strings
	names := []string{"Tony", "Tamara", "Warren", "Arnold", "Leon", "Peterson", "Brandon"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Tony"))
}
