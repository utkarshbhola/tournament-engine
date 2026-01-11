package engine

import (
	"tournament-engine/internal/model"
	"tournament-engine/internal/util"
)

type FightResult struct {
	Winner model.Fighter
}

func runRound(fighters []model.Fighter) []model.Fighter {
	results := make(chan FightResult, len(fighters)/2)

	for i:=0;i<len(fighters); i += 2{
		a:= fighters[i]
		b:= fighters[i+1]

		go func() {
			rng := util.NewRng()
			winner := simulateFight(a,b,rng)
			results <- FightResult{Winner:winner}
		}()
	}
	var winners []model.Fighter
	for i:= 0;i < len(fighters)/2;i++{
		res := <-results
		winners = append(winners,res.Winner)
	}
	return winners
}