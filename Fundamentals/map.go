package main

import (
	"fmt"
)

func main() {
	obj := map[int]string{
		1: "Brendo",
		2: "Maria",
		3: "João",
	}

	obj[4] = "Julia"
	delete(obj, 3)

	for index, name := range obj {
		fmt.Printf("%v = %v \n", index, name)
	}
}
