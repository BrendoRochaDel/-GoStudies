package main

import "fmt"

func age(year int) int {
	x := 0

	for year < 2022 {
		x++
		year++
	}

	return x
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, "")
	}

	fmt.Println("")

	x := 0
	for x < 10 {
		fmt.Print(x, "")
		x++
	}

	year := 1998
	fmt.Println("")
	fmt.Println(age(year), "years")
}
