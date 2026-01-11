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

func printFighters(fighters []model.Fighter) {
	fmt.Println("AVAILABLE FIGHTERS:")
	fmt.Println("ID  NAME     HT  HP  ATK  TYPE")

	for i, f := range fighters {
		fmt.Printf(
			"%d   %-8s %3d %3d %3d  %s\n",
			i, f.Name, f.HeightCM, f.MaxHealth, f.Attack, f.Archetype,
		)
	}
	fmt.Println()
}

func chooseFighter(fighters []model.Fighter) int {
	var choice int

	fmt.Print("Choose your fighter (enter ID): ")
	fmt.Scan(&choice)

	if choice < 0 || choice >= len(fighters) {
		fmt.Println("Invalid choice. Defaulting to fighter 0.")
		return 0
	}

	return choice
}

func main() {
	printIntroStory()

	fighters := loadFighters()
	printFighters(fighters)

	playerIndex := chooseFighter(fighters)
	playerName := fighters[playerIndex].Name

	fmt.Println("\nYou chose:", playerName)
	fmt.Println("-----------------------------------")

	champion := engine.RunTournament(fighters)

	fmt.Println()
	if champion.Name == playerName {
		fmt.Println("YOU WON THE TOURNAMENT!")
	} else {
		fmt.Println("AND THE CHAMPION IS !!", champion.Name)
		fmt.Println("You were eliminated earlier.")
	}
}
