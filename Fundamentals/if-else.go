package main

import "fmt"

func main() {
	year := 20

	if year >= 18 {
		fmt.Println("Are you able to drive!")
	} else {
		fmt.Println("You are not fit to drive!")
	}
}
