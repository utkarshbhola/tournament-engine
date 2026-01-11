package model

type SpecialMove struct {
	Name        string
	TriggerProb float64
	DamageBoost int
}

type Fighter struct {
	Name      string
	HeightCM  int
	MaxHealth int
	Health    int
	Attack    int
	Archetype Archetype
	Special   SpecialMove
}
