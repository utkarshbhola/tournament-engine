package engine

import (
	"math/rand"
	"tournament-engine/internal/model"
)

func applyUpgrade(f *model.Fighter, rng *rand.Rand){
	if rngIntn(2) == 0{
		f.MaxHealth += 10
		f.Health = f.MaxHealth
	}
	else{
		f.Attack += 2
	}
}