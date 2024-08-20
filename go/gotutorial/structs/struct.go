package main

import "fmt"

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

func Super(s *Saiyan) {
	s.Power += 10000
}
func NewSaiyan(name string, power int, father *Saiyan) *Saiyan {
	return &Saiyan{
		Name:   name,
		Power:  power,
		Father: father,
	}
}
func main() {
	goku1 := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}
	goku2 := &Saiyan{"Musk", 8000, &goku1}
	goku3 := NewSaiyan("Gohan", 9001, goku2)
	Super(goku2)
	Super(goku3)
	fmt.Println("goku1", goku1.Name, goku1.Power, goku1.Father)
	fmt.Println("goku2", goku2.Name, goku2.Power, goku2.Father)
	fmt.Println("goku3", goku3.Name, goku3.Power, goku3.Father)
}
