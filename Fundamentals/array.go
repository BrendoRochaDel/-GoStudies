package main

import "fmt"

func main() {
	array := [5]string{
		"Brendo",
		"Juliana",
		"João",
		"Thiago",
		"Matheus",
	}

	array[4] = "Maria"

	for indice, name := range array {
		fmt.Printf("%v = %v \n", indice, name)
	}
}
