package models

// Vallis : Represents Vallis (Fortuna) cycle status
type Vallis struct {
	ID       string `json:"id"`
	Expiry   string `json:"expiry"`
	TimeLeft string `json:"timeLeft"`
	IsWarm   bool   `json:"isWarm"`
}

//GetCurrentState : Show current Vallis state
func (vallis Vallis) GetCurrentState() string {
	if vallis.IsWarm {
		return "Warm \U0001F525"
	} else {
		return "Cold \U00002744"
	}
}
