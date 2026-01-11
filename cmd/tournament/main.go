package main

import (
	"fmt"
	"tournament-engine/internal/engine"
	"tournament-engine/internal/model"
)

func loadFighters() []model.Fighter {
	return []model.Fighter{
		{"Rohit", 178, 120, 120, 14, model.Power,
			model.SpecialMove{"Heavy Elbow", 0.20, 6}},
		{"Aman", 172, 105, 105, 16, model.Speed,
			model.SpecialMove{"Rapid Combo", 0.25, 5}},
		{"Suresh", 180, 115, 115, 15, model.Technique,
			model.SpecialMove{"Joint Lock", 0.22, 5}},
		{"Vikas", 175, 110, 110, 14, model.Power,
			model.SpecialMove{"Overhand Smash", 0.18, 7}},
		{"Deepak", 170, 100, 100, 17, model.Speed,
			model.SpecialMove{"Counter Strike", 0.28, 4}},
		{"Ravi", 176, 112, 112, 15, model.Wildcard,
			model.SpecialMove{"Dirty Clinch", 0.15, 6}},
	}
}

func main() {
	fighters := loadFighters()
	champion := engine.RunTournament(fighters)

	fmt.Println("AND THE CHAMPION IS !!", champion.Name)
}
