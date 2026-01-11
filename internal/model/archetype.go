package model

type Archetype int

const (
	Power Archetype = iota
	Speed
	Technique
	Wildcard
)

func (a Archetype) String() string {
	switch a {
	case Power:
		return "Power"
	case Speed:
		return "Speed"
	case Technique:
		return "Technique"
	case Wildcard:
		return "Wildcard"
	default:
		return "Unknown"
	}
}
