package main

import "fmt"

func main() {
	menu := map[string]float64{
		"Ugali": 70.0,
		"Mboga": 30.0,
		"Mayai": 40.0,
		"Rice":  30.0,
	}

	//fmt.Println(menu)
	//fmt.Println(menu["Ugali"])

	//looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

}
