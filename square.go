package main

import (
	"math"

	"github.com/thats-insane/quiz1/round"
)

func Square(side float64) (float64, float64) {
	return round.RoundFloat(side*4, 10), round.RoundFloat(math.Pow(side, 2), 10)
}
