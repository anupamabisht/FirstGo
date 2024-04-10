package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	RollNumber int
	Name       string
	Country    string
}
type RollNumber struct {
	Max int
	Min int
}

type EmployeeRange struct {
	Start string
}

type Country struct {
	CountryCode string
}

func main() {

	employees := []Employee{
		{RollNumber: 1, Name: "Alice", Country: "Sweden"},
		{RollNumber: 2, Name: "James", Country: "Finland"},
		{RollNumber: 3, Name: "Nitara", Country: "Denmark"},
		{RollNumber: 4, Name: "Josh", Country: "Norway"},
		{RollNumber: 5, Name: "Alan", Country: "France"},
		{RollNumber: 6, Name: "Jamie", Country: "Iceland"},
		{RollNumber: 7, Name: "Mike", Country: "Germany"},
		{RollNumber: 8, Name: "John", Country: "Netherland"},
		{RollNumber: 9, Name: "Noah", Country: "Hungary"},
		{RollNumber: 10, Name: "Balmain", Country: "Czech"},
		{RollNumber: 11, Name: "Joe", Country: "Swizerland"},
		{RollNumber: 12, Name: "Anna", Country: "Spain"},
		{RollNumber: 13, Name: "Bengt", Country: "Greece"},
		{RollNumber: 14, Name: "Bartek", Country: "Greenland"},
		{RollNumber: 15, Name: "Meesho", Country: "Sweden"},
		{RollNumber: 16, Name: "Jakob", Country: "Italy"},
	}
	fmt.Println(employees)

	rollnumber := []RollNumber{
		{Min: 1, Max: 10}}

	fmt.Println(rollnumber)

	startletter := []EmployeeRange{
		{Start: "A"}}

	fmt.Println(startletter)

	country := []Country{
		{CountryCode: "S"},
	}
	fmt.Println(country)

	filteredEmployee := getFilteredValue(employees, rollnumber, startletter, country)
	fmt.Println("Filtered Employees:")
	for _, rollnumberange := range filteredEmployee {
		fmt.Printf("RollNumber: %d, Name: %s, Country: %s\n", rollnumberange.RollNumber, rollnumberange.Name, rollnumberange.Country)

	}
}

func getFilteredValue(employees []Employee, rollnumber []RollNumber, starletter []EmployeeRange, country []Country) []Employee {
	var filteredEmployee []Employee

	for _, rollnumberrange := range employees {
		for _, roll := range rollnumber {
			for _, emprange := range starletter {
				for _, countycode := range country {
					{
						if (rollnumberrange.RollNumber >= roll.Min && rollnumberrange.RollNumber <= roll.Max) &&
							strings.HasPrefix(rollnumberrange.Name, emprange.Start) && strings.HasPrefix(rollnumberrange.Country, countycode.CountryCode) {
							filteredEmployee = append(filteredEmployee, rollnumberrange)
						}

					}
				}
			}
		}

	}
	return filteredEmployee
}
