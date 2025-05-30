package mathutils

import (
	"errors"

	"github.com/niketa1493/solid-spork/pkg/internal"
)

func Add(a, b float64) string {
	result := a + b
	return internal.PrintResult(result)
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
