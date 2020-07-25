package main

import "fmt"

func main() {
	//Strings, which can be added together with +.
	fmt.Println("luiz" + "ricardo")

	//Integers and floats.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	//Booleans, with boolean operators as you’d expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
