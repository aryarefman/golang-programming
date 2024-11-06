package main

import "testing"

// Pengujian fungsi add
func TestAdd(t *testing.T) {
	result := add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5 but got %d", result)
	}
}

// Pengujian fungsi swap
func TestSwap(t *testing.T) {
	a, b := swap("first", "second")
	if a != "second" || b != "first" {
		t.Errorf("Expected (second, first) but got (%s, %s)", a, b)
	}
}

// Pengujian fungsi divide
func TestDivide(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Errorf("Expected error for division by zero but got nil")
	}

	result, err := divide(10, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("Expected 5 but got %d", result)
	}
}
