package main

import "testing"

// TestGet tests the Get function
func TestGet(t *testing.T) {
	m := NewDefaultDict(0)
	if m.Get(100) != 0 {
		t.Errorf("Get working incorrectly with int keys")
	}
	if m.Get("aaaa") != 0 {
		t.Errorf("Get working incorrectly with string keys")
	}
}

// TestSet tests the Set function
func TestSet(t *testing.T) {
	m := NewDefaultDict("")
	err := m.Set("name", "jeff")
	if err != nil {
		t.Errorf("Set working incorrectly with string keys and string values")
	}
}
