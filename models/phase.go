package models

type Phase struct {
	Selections []Selection `json:"selections"`
}

func NewPhase(selections []Selection) Phase {
	return Phase {
		Selections: selections,
	}
}