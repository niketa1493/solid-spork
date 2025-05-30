package mathutils

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(2, 3) != "The result is: 5.00" {
		t.Error("Add failed")
	}
}

func TestSubtract(t *testing.T) {
	if Subtract(5, 3) != 2 {
		t.Error("Subtract failed")
	}
}

func TestMultiply(t *testing.T) {
	if Multiply(4, 2.5) != 10 {
		t.Error("Multiply failed")
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil || result != 5 {
		t.Error("Divide failed")
	}

	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Expected error when dividing by zero")
	}
}
