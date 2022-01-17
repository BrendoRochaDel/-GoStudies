package main

import "fmt"

func main() {
	slice := []string{
		"Brendo",
		"Juliana",
		"João",
		"Thiago",
		"Matheus",
	}

	// Add Maria
	slice = append(slice, "Maria")

	//Remove João
	slice = append(slice[0:2], slice[3:len(slice)]...)

	for indice, name := range slice {
		fmt.Printf("%v = %v \n", indice, name)
	}
}
