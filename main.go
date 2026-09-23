package main

import "fmt"

type Person interface {
	Who() string
}

type Human struct{}

type Robot struct{}

func (h Human) Who() string {
	return "Привет, я человек"
}

func (r Robot) Who() string {
	return "Привет, я робот"
}

func main() {
	var p Person

	p = Robot{}
	fmt.Println(p.Who())

	p = Human{}
	fmt.Println(p.Who())
}