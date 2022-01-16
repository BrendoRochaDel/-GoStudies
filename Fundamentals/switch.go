package main

import (
	"fmt"
)

var x interface{}

func main() {
	x = 45

	switch x.(type) {
	case int:
		fmt.Println("Integer type")
	case bool:
		fmt.Println("Boolean type")
	case string:
		fmt.Println("String type")
	case float64:
		fmt.Println("Float type")
	default:
		fmt.Println("Unidentified type")
	}
}
