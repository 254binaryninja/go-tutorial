package main

import "fmt"

func main() {
	age := 20
	name := "Musandu"

	//Printf (formated strings) %_ =format specifier
	fmt.Printf("My age is %v and my name is %v \n", age, name)
	fmt.Printf("My age is %v and my name is %q \n", age, name)

	var str = fmt.Sprintf("My age is %v and my name is %q \n", age, name)
	fmt.Println("Here is a save string :", str)
}
