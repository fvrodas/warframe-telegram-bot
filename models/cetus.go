package models

// Cetus : Repesents Cetus cycle status
type Cetus struct {
	ID          string `json:"id"`
	Expiry      string `json:"expiry"`
	Activation  string `json:"activation"`
	IsDay       bool   `json:"isDay"`
	State       string `json:"state"`
	TimeLeft    string `json:"timeLeft"`
	IsCetus     bool   `json:"isCetus"`
	ShortString string `json:"shortString"`
}

//GetCurrentState : Show current Vallis state
func (cetus Cetus) GetCurrentState() string {
	if cetus.IsDay {
		return "Day \U0001F31E"
	}
	return "Night \U0001F31A"
}
