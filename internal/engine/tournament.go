package engine

import "tournament-engine/internal/model"

func RunTournament(fighters []model.Fighter) model.Fighter {
	for len(fighters) > 1{
		fighters = runRound(fighters)
	}
	return fighters[0]
}