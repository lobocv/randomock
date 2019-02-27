package main

import (
	"fmt"
	"testing"

	"github.com/lobocv/randomock"
)

func TestDice(t *testing.T) {

	t.Run("Only One Return Value", func(t *testing.T) {
		// Expect the result to be 4
		expected := 4
		rand := randomock.NewRandoMock().Add("roll", float64(expected))

		dice := NewDice(6, rand)

		for roll := 0; roll < 10; roll++ {
			result := dice.Roll()
			if result != expected {
				t.Fatalf("expected to roll %d, got %d", expected, result)
			}

		}
	})

	t.Run("Many Different Return Values", func(t *testing.T) {
		// Expect the result to be 4
		expectedList := []float64{4, 2, 3, 1}
		rand := randomock.NewRandoMock().Add("roll", expectedList...)

		dice := NewDice(6, rand)

		// Note we are only calling as many times as we have provided return values
		// If we called more than that, we would get a panic due to the default Randomock policy
		for roll := 0; roll < len(expectedList); roll++ {
			result := dice.Roll()
			expected := int(expectedList[roll])
			if result != expected {
				t.Fatalf("expected to roll %d, got %d", expected, result)
			}
		}
	})

	t.Run("More calls than Return Values", func(t *testing.T) {
		defer func() {
			r := recover()
			fmt.Printf("We got a panic, as expected! Panic: %v\n", r)
		}()

		// Expect the result to be 4
		expectedList := []float64{4, 2, 3, 1}
		rand := randomock.NewRandoMock().Add("roll", expectedList...)

		dice := NewDice(6, rand)

		for roll := 0; roll < 10; roll++ {
			result := dice.Roll()
			expected := int(expectedList[roll])
			if result != expected {
				t.Fatalf("expected to roll %d, got %d", expected, result)
			}
		}
	})

	t.Run("Different Randomock Policy", func(t *testing.T) {
		// Expect the result to be 4
		expectedList := []float64{4, 2, 3, 1}
		randomock.SetDefaultPolicy(randomock.WrapAroundPolicy)
		rand := randomock.NewRandoMock().Add("roll", expectedList...)

		dice := NewDice(6, rand)

		for roll := 0; roll < 10; roll++ {
			result := dice.Roll()
			expected := int(expectedList[roll%len(expectedList)]) // Note we wrap around here
			if result != expected {
				t.Fatalf("expected to roll %d, got %d", expected, result)
			}
		}
	})

}
