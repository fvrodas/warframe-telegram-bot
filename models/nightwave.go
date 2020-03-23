package models

// Nightwave : Represents information about the Nigthwave
type Nightwave struct {
	ID                 string      `json:"id"`
	Activation         string      `json:"activation"`
	Expiry             string      `json:"expiry"`
	RewardTypes        []string    `json:"rewardTypes"`
	Season             int         `json:"season"`
	Tag                string      `json:"tag"`
	Phase              int         `json:"phase"`
	PossibleChallenges []Challenge `json:"posibleChallenges"`
	ActiveChallenges   []Challenge `json:"activeChallenges"`
}

// Challenge : Challenge details
type Challenge struct {
	ID         string `json:"id"`
	Activation string `json:"activation"`
	Expiry     string `json:"expiry"`
	IsDaily    bool   `json:"isDaily"`
	IsElite    bool   `json:"isElite"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Reputation int    `json:"reputation"`
}
