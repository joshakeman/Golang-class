package main

import "fmt"

func main() {

	z := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Josh",
		friends: map[string]int{
			"guy": 22,
			"g":   55,
			"s":   333,
		},
		favDrinks: []string{"coke", "water"},
	}

	fmt.Println(z)
}
