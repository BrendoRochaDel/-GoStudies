package main

import "fmt"

// var x int = 42
// var y string = "James Bond"
// var z bool = true

type myType int

var my myType
var myAux int

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%v\n%v\n%v\n", x, y, z)

	my = 42
	fmt.Println(my)
	myAux = int(my)
	fmt.Printf("%v %T \n", myAux, myAux)
}
