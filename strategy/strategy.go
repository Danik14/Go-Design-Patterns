package main

import "fmt"

func main() {
	//as you can see i created 1 variable instrument
	//but i can switch different strategies
	//without adding new variables
	instrument := NewInstrument(Hammer{})
	instrument.PerformAction()
	instrument.SetInstrument(Screwdriver{})
	instrument.PerformAction()
	instrument.SetInstrument(Pickaxe{})
	instrument.PerformAction()
}

type Instrument interface {
	Action()
}

type Hammer struct{}

func (hammer Hammer) Action() {
	fmt.Println("I am used to hammer nails")
}

type Screwdriver struct{}

func (scr Screwdriver) Action() {
	fmt.Println("I am used to tighten screws")
}

type Pickaxe struct{}

func (p Pickaxe) Action() {
	fmt.Println("I am used to dig in mines")
}

type ConcreteInstument struct {
	instrument Instrument
}

func NewInstrument(i Instrument) *ConcreteInstument {
	return &ConcreteInstument{
		instrument: i,
	}
}

func (t *ConcreteInstument) PerformAction() {
	t.instrument.Action()
}

func (t *ConcreteInstument) SetInstrument(i Instrument) {
	t.instrument = i
}
