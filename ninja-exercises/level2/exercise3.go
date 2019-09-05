package main

import "fmt"

func main() {

	x := map[string]int{
		"Guy":       44,
		"Other guy": 2,
	}

	fmt.Println(x["Guy"])
}
