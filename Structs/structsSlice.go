package main

import "fmt"

type People struct {
	name   string
	age    int
	active bool
}

func main() {
	peopleList := []People{}
	peopleAux := People{"Brendo", 23, true}

	peopleList = append(peopleList, peopleAux)

	peopleAux = People{"Maria", 19, false}
	peopleList = append(peopleList, peopleAux)

	for _, people := range peopleList {
		fmt.Printf("nome: %v \nidade: %v\n", people.name, people.age)
		fmt.Println("=====================================")
	}
}
