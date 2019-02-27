package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lobocv/randomock"
)

// Dice is a dice that you can roll
type Dice struct {
	sides int
	rand  randomock.Randomizer
}

// NewDice creates a new dice with the given number of sides.
func NewDice(sides int, rand randomock.Randomizer) *Dice {
	return &Dice{sides: sides, rand: rand}
}

// Roll simulates rolling a dice
func (d *Dice) Roll() int {
	// We named this call "roll" so we can easily mock the specific call out in our tests
	return d.rand.Intn("roll", d.sides)
}

// This is a simple example of rolling a dice.
// Since this is not test code, use a real random number generator
func main() {
	rand.Seed(time.Now().Unix())
	dice := NewDice(6, &randomock.Random{})

	for roll := 0; roll < 10; roll++ {
		result := dice.Roll()
		fmt.Printf("Rolling dice attempt %d.... got %d\n", roll, result)
	}
}
