package main

import (
	"testing"
)

func TestSetAndGetAddressFails(t *testing.T) {
	t.Log("Input: Ini Bumiku A.20 RT/RW 20/21 @Hallo")
	t.Log("Expected: Ini Bumiku A.20 RT/RW 20/21 @Hallo")

	var e Employee
	e.SetAddress("Ini Bumiku A.20 RT/RW 20/21 @Hallo")
	e.GetAddress()
	if e.Address != "Ini Bumiku A.20 RT/RW 20/21 @Hallo" {
		t.Error("Sorry, not accept symbol")
	}
}

func TestSetAndGetAddressPass(t *testing.T) {
	t.Log("Input: Ini Bumiku A.20 RT/RW 20/21")
	t.Log("Expected: Ini Bumiku A.20 RT/RW 20/21")

	var e Employee
	e.SetAddress("Ini Bumiku A.20 RT/RW 20/21")
	e.GetAddress()
	if e.Address != "Ini Bumiku A.20 RT/RW 20/21" {
		t.Error("Sorry, not accept symbol")
	}
}
