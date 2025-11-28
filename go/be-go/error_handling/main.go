package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImpl  = errors.New("Not implemented")
	ErrNotFound = errors.New("Not Found")
)

// cant := outside fn at package level
var print = fmt.Printf

type truckType interface {
	LoadCargo() error
	UnloadCargo() error
}

// first letter is capital means this struct can beimported outside this package (same logic for var names)
type Truck struct {
	id    string
	cargo int
}

func (t *Truck) LoadCargo() error {
	return nil
}

func (t *Truck) UnloadCargo() error {
	return nil
}

func (t *ElectricTruck) LoadCargo() error {
	return nil
}

func (t *ElectricTruck) UnloadCargo() error {
	return nil
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func process(truck truckType) error {
	// fmt.Printf("Processing Truck: %s\n", truck.id) // interface dont define fields

	err := truck.LoadCargo()
	if err != nil {
		return fmt.Errorf("error loading %w\n", err)
	}

	err = truck.UnloadCargo()
	if err != nil {
		a := fmt.Errorf("error unloading %w\n", err)
		return a
	}

	return nil
}

func main() {
	trucks := []*Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}
	//does this underthe hood
	/*
		trucks := []*Truck{
			&Truck{id: "Truck-1"},
			&Truck{id: "Truck-2"},
			&Truck{id: "Truck-3"},
		}
	*/

	eTrucks := []*ElectricTruck{
		{id: "E-Truck-1"},
	}

	for _, truck := range trucks {
		print("Truck %s arrived\n", truck.id)

		err := process(truck)
		if err != nil {
			log.Fatalf("Error processing Truck: %s\n", err)
		}

		print("Truck %s departed\n", truck.id)
	}

	for _, truck := range eTrucks {
		print("Truck %s arrived\n", truck.id)

		err := process(truck)
		if err != nil {
			log.Fatalf("Error processing Truck: %s\n", err)
		}

		print("Truck %s departed\n", truck.id)
	}
}
