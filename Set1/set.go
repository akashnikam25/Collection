package main

import (
	"fmt"
	"os"
)

type animalSet map[string]struct{}

func (s animalSet) add(animal string) {
	if !s.has(animal) {
		s[animal] = struct{}{}
	} else {
		fmt.Println("animal already exist")
	}
}
func (s animalSet) remove(animal string) {
	if s.has(animal) {
		delete(s, animal)
	} else {
		fmt.Println("animal does not exist")
	}

}
func (s animalSet) has(animal string) (ok bool) {
	_, ok = s[animal]
	return
}
func (s animalSet) print() {
	if len(s) != 0 {
		for key, _ := range s {
			fmt.Print(key, " ")
		}
		fmt.Println()
	} else {
		fmt.Println("animal set is empty")
	}
}

func enterChoice() (choice int) {
	fmt.Println("0.Exit")
	fmt.Println("1.Add")
	fmt.Println("2.Remove")
	fmt.Println("3.Has")
	fmt.Println("4.Print")
	fmt.Println("Enter choice")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return
}
func enterAnimal() (animal string) {
	_, err := fmt.Scan(&animal)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return
}
func main() {
	s := animalSet{}
	choice := 1 //by default 1 value we set it

	for choice != 0 {
		switch choice {
		case 1:
			fmt.Println("Enter animal to add")
			s.add(enterAnimal())
		case 2:
			fmt.Println("Enter animal to Remove")
			s.remove(enterAnimal())
		case 3:
			fmt.Println("Enter animal to check exist")
			s.has(enterAnimal())
		case 4:
			fmt.Println("Print all set value")
			s.print()
		}
		choice = enterChoice()
		fmt.Println("choice", choice)
	}
}
