package main

import (
	"math"

	"github.com/thats-insane/quiz1/round"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return round.RoundFloat(math.Pow(c.Radius, 2)*3.14, 10)
}

func (c Circle) Perimeter() float64 {
	return round.RoundFloat(c.Radius*2*3.14, 10)
}
