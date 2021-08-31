package main

import "fmt"

type animal interface {
	makeSound() string
}

type dog struct {
	name string
}

type cat struct {
	name string
}

func (d dog) makeSound() string {
	return "Wao Wao"
}

func (c cat) makeSound() string {
	return "Miau"
}

func animalMakeSound(a animal) {
	fmt.Println(a.makeSound())
}

func main() {
	var cafu dog = dog{name: "Lelouch"}
	var sisa cat = cat{name: "Suzaku"}

	animalMakeSound(cafu)
	animalMakeSound(sisa)

}
