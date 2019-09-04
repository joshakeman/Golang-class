package main

import "fmt"

type hotdog int

var x hotdog

func main() {
	fmt.Printf("%v\t%T\t", x, x)
}
