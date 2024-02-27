package golib

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(3, 4)
	if result != 7 {
		t.Errorf("Add(3, 4) expected 7, but got %d", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(7, 3)
	if result != 4 {
		t.Errorf("Subtract(7, 3) expected 4, but got %d", result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(3, 4)
	if result != 12 {
		t.Errorf("Multiply(3, 4) expected 12, but got %d", result)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil || result != 5 {
		t.Errorf("Divide(10, 2) expected 5, but got %d with error: %v", result, err)
	}

	_, err = Divide(10, 0)
	if err == nil {
		t.Errorf("Divide(10, 0) expected an error, but got nil")
	}
}
