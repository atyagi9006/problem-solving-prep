package main

import "fmt"

type Fowl interface {
	Lay() FowlEgg
}

type FowlCreator func() Fowl

type FowlEgg struct {
	CreateFowl FowlCreator
}

func (egg *FowlEgg) Hatch() Fowl {

	return egg.CreateFowl()
}

type Goose struct {
	ParentId string
}

func (g Goose) Lay() FowlEgg {

	cf := func()Fowl{
		return Goose{
			ParentId: "mine",
		}
	}

	egg := FowlEgg{
		CreateFowl: cf,
	}

	return egg
}

func NewGoose() Goose {
	return Goose{}
}

func main() {
	var goose1 Goose
	var egg FowlEgg = goose1.Lay()
	var childGoose Fowl = egg.Hatch()
	fmt.Println(childGoose)
}
