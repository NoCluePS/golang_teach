package main

import "fmt"

type Player struct {
	Health int
	MaxHealth int
	Energy int
	MaxEnergy int
	Name string
}

func (player *Player) logPlayerData() {
	fmt.Printf("\nName: %s\nHealth: %d\nEnergy: %d\n", player.Name, player.Health, player.Energy)
}

func (player *Player) addHealth(health int) {
	neededHealth := player.Health + health

	if neededHealth > player.MaxHealth {
		player.Health = player.MaxHealth
	} else {
		player.Health = neededHealth 
	}
}

func (player *Player) removeHealth(health int) {
	neededHealth := player.Health - health

	if neededHealth < 0 {
		player.Health = 0
	} else {
		player.Health = neededHealth 
	}
}


func (player *Player) addEnergy(energy int) {
	neededEnergy := player.Energy + energy

	if neededEnergy > player.MaxEnergy {
		player.Energy = player.MaxEnergy
	} else {
		player.Energy = neededEnergy 
	}
}

func (player *Player) removeEnergy(energy int) {
	neededEnergy := player.Energy - energy

	if neededEnergy < 0 {
		player.Energy = 0
	} else {
		player.Energy = neededEnergy 
	}
}

func main() {
	player := Player{
		Health: 12,
		MaxHealth: 100,
		Energy: 23,
		MaxEnergy: 1000,
		Name: "Thomas",
	}
	player.addEnergy(32)
	player.logPlayerData()
	player.addHealth(12)
	player.logPlayerData()
	player.removeEnergy(2000)
	player.logPlayerData()
	player.removeHealth(2000)
	player.logPlayerData()

}
