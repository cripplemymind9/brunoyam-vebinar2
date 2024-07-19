package main

import (
	"fmt"
)

type Vehicle interface {
	Start()
	Stop()
}

type Car struct {
	Mark string
	Model string
	Engine float64
}

type Moto struct {
	Mark string
	Model string
	Engine float64
}

func (c Car) Start() {
	fmt.Printf("%v %v %v started!\n", c.Mark, c.Model, c.Engine)
}

func (c Car) Stop() {
	fmt.Printf("%v %v %v stopped!\n", c.Mark, c.Model, c.Engine)
}

func (m Moto) Start() {
	fmt.Printf("%v %v %v started!\n", m.Mark, m.Model, m.Engine)
}

func (m Moto) Stop() {
	fmt.Printf("%v %v %v stopped!\n", m.Mark, m.Model, m.Engine)
}

func main() {
	car := Car {
		Mark: "BMW",
		Model: "X5",
		Engine: 3.0,
	}

	moto := Moto {
		Mark: "Honda",
		Model: "CBR650R",
		Engine: 3.0,
	}
	
	car.Start()
	moto.Start()
	car.Stop()
	moto.Stop()
}