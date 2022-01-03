package main

import (
	"fmt"
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
	var regex, err = regexp.Compile(`(?i)[a-z]+`)
	if err != nil {
		fmt.Println(err.Error())
	}
	//Debugging
	e.Name = regex.FindString(e.Name)
	// fmt.Printf("Value of e.Name after Regex %+v", e.Address)
	return e.Name
}

func (e *Employee) GetName() string {
	return e.Name
}

//Address
func (e *Employee) SetAddress(address string) string {
	e.Address = address
	var regex, err = regexp.Compile(`^[a-zA-Z0-9_.,-/ ]*$`)
	if err != nil {
		fmt.Println(err.Error())
	}

	e.Address = regex.FindString(e.Address)
	//fmt.Printf("Value of e.Address after Regex %+v", e.Address)
	return e.Address
}

func (e *Employee) GetAddress() string {
	return e.Address
}

// Age
func (e *Employee) SetAge(age string) int {
	ageInt, _ := strconv.Atoi(age)
	e.Age = ageInt
	var regex, err = regexp.Compile(`\d`)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Debugging
	AgeStr := strconv.Itoa(e.Age)
	AgeStr = regex.FindString(AgeStr)
	//fmt.Printf("Value of e.Age/AgeStr after Regex %+v", AgeStr)

	return e.Age
}

func (e *Employee) GetAge() int {
	return e.Age
}
