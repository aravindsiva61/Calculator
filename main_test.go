package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(5, 3)
	expected := 8
	if result != expected {
		t.Errorf("add = %d; expected %d", result, expected)
	}
}

func TestSub(t *testing.T) {
	result := sub(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("sub = %d; expected %d", result, expected)
	}
}

func TestMul(t *testing.T) {
	result := mul(5, 3)
	expected := 15
	if result != expected {
		t.Errorf("mul = %d; expected %d", result, expected)
	}
}
