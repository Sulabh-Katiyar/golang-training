package main

import "fmt"

func main() {

	type Weather struct {
		city string
		temp float64
	}

	var m float64 = map[string]Weather{
		"Delhi": 31.45,
		"UP":    35.10,
	}
	fmt.Println(m["Delhi"])
}
