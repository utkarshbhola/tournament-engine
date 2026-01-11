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

	var winners []model.Fighter

	limit := len(fighters)

	// If odd, last fighter gets a bye
	if limit%2 == 1 {
		winners = append(winners, fighters[limit-1])
		limit-- // exclude last fighter from fights
	}

	for i := 0; i < limit; i += 2 {
		a := fighters[i]
		b := fighters[i+1]

		go func() {
			rng := util.NewRng()
			winner := simulateFight(a, b, rng)
			results <- FightResult{Winner: winner}
		}()
	}

	for i := 0; i < limit/2; i++ {
		res := <-results
		winners = append(winners, res.Winner)
	}

	return winners
}
