package main

import "fmt"

type People struct {
	name   string
	age    int
	active bool
}

type Phone struct {
	number string
	people People
}

func main() {
	phone1 := Phone{
		number: "+55 (11)4002-8922",
		people: People{
			name:   "Maria",
			age:    27,
			active: true,
		},
	}

	fmt.Println(phone1)
}
