package engine

import (
	"math/rand"
	"tournament-engine/internal/model"
)

var advantage = map[model.Archetype]model.Archetype{
	model.Power:     model.Speed,
	model.Speed:     model.Technique,
	model.Technique: model.Power,
}

func hasAdvantage(a, b model.Archetype) bool {
	return advantage[a] == b
}

func simulateFight(a, b model.Fighter, rng *rand.Rand) model.Fighter {
	a.Health = a.MaxHealth
	b.Health = b.MaxHealth

	attacker, defender := &a, &b

	for attacker.Health > 0 && defender.Health > 0 {
		damage := attacker.Attack

		if hasAdvantage(attacker.Archetype,defender.Archetype){
			damage += 3
		}
		
		if rng.Float64() < attacker.Special.TriggerProb {
			damage += attacker.Special.DamageBoost
		}

		defender.Health -= damage
		attacker,defender = defender, attacker
	}

	if a.Health > 0{
		return a
	}
	return b
}
