package main

import "fmt"

type People struct {
	name   string
	age    int
	active bool
}

func main() {
	people1 := People{
		name:   "João",
		age:    18,
		active: true,
	}

	people2 := People{"Brendo", 23, false}

	fmt.Println(people1)
	fmt.Println(people2)
}
