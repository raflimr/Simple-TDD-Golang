package main

import (
	"testing"
)

func TestSetAndGetAgeFails(t *testing.T) {
	t.Log("Input: 10,8")
	t.Log("Expected: 11")

	var e Employee
	e.SetAge("10,8")
	e.GetAge()
	if e.Age != 11 {
		t.Error("Cannot be rounded, only accepts integers")
	}
}

func TestSetAndGetAgePass(t *testing.T) {
	t.Log("Input: 10")
	t.Log("Expected: 10")

	var e Employee
	e.SetAge("10")
	e.GetAge()
	if e.Age != 10 {
		t.Error("Cannot be rounded, only accepts integers")
	}
}
