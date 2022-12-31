//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests
package main

import "testing"

func TestHealth(t *testing.T) {
	p := Player{
		Health: 10,
		MaxHealth: 100,
		Energy: 20,
		MaxEnergy: 150,
		Name: "Pijus",
	}

	p.addHealth(1000)
	if p.Health > p.MaxHealth {
		t.Errorf("Health can only be at max: %v, want: %v", p.Health, p.MaxHealth)
	}

	p.removeHealth(50)
	if p.Health != 50 {
		t.Errorf("Didn't remove the amount of health wanted. Got: %v, want: %v", p.Health, 100)
	}

	p.removeHealth(10000)
	if p.Health < 0 {
		t.Errorf("Health can't go below 0")
	}
}

func TestEnergy(t *testing.T) {
	p := Player{
		Health: 10,
		MaxHealth: 100,
		Energy: 20,
		MaxEnergy: 150,
		Name: "Pijus",
	}

	p.addEnergy(1000)
	if p.Energy > p.MaxEnergy {
		t.Errorf("Energy can only be at max: %v, want: %v", p.Energy, p.MaxEnergy)
	}

	p.removeEnergy(50)
	if p.Energy != 100 {
		t.Errorf("Didn't remove the amount of energy wanted. Got: %v, want: %v", p.Energy, 100)
	}

	p.removeEnergy(10000)
	if p.Energy < 0 {
		t.Errorf("Energy can't go below 0")
	}
}