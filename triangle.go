package main

import (
	"math"

	"github.com/thats-insane/quiz1/round"
)

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return round.RoundFloat((t.Base*t.Height)/2, 10)
}

func (t Triangle) Perimeter() float64 {
	return round.RoundFloat(t.Base+t.Height+(math.Sqrt(math.Pow(t.Base, 2)+math.Pow(t.Height, 2))), 10)
}
