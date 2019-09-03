package main

import "fmt"

func main() {
	x := "Jimbo bob"
	y := 42
	z := true

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
