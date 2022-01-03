package main

import (
	"regexp"
	"strconv"
)

type EmployeeInterface interface {
	SetName(name string) string
	SetAddress(address string) string
	SetAge(age string) int
	GetName() string
	GetAddress() string
	GetAge() int
}

type Employee struct {
	Name, Address string
	Age           int
}

// Name
func (e *Employee) SetName(name string) string {
	e.Name = name
	var regex, _ = regexp.Compile(`(?i)[a-z]+`)
	return regex.FindString(e.Name)
}

func (e *Employee) GetName() string {
	return e.Name
}

//Address
func (e *Employee) SetAddress(address string) string {
	e.Address = address
	var regex, _ = regexp.Compile(`^[a-zA-Z0-9_.,-/ ]*$`)
	return regex.FindString(e.Address)
}

func (e *Employee) GetAddress() string {
	return e.Address
}

// Age
func (e *Employee) SetAge(age string) int {
	ageInt, _ := strconv.Atoi(age)
	e.Age = ageInt
	var regex, _ = regexp.Compile(`\d`)
	AgeStr := strconv.Itoa(e.Age)
	AgeStr = regex.FindString(AgeStr)
	return e.Age
}

func (e *Employee) GetAge() int {
	return e.Age
}
