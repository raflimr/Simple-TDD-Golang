package main

import (
	"testing"
)

func TestSetAndGetNameFails(t *testing.T) {
	t.Log("Input: Rafli123")
	t.Log("Expected: Rafli")

	var e Employee
	e.SetName("Rafli123")
	e.GetName()
	if e.Name != "Rafli123" {
		t.Error("Sorry, input only accepts letters from a-z")
	}
}

func TestSetAndGetNamePass(t *testing.T) {
	t.Log("Input: Rafli")
	t.Log("Expected: Rafli")

	var e Employee
	e.SetName("Rafli")
	e.GetName()
	if e.Name != "Rafli" {
		t.Error("Wrong, Should only write letters")
	}

}
