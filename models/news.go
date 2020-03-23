package models

import (
	"fmt"
	"log"
	"time"
)

// News : News about Warframe
type News struct {
	Date        string `json:"date"`
	Imagelink   string `json:"imageLink"`
	PrimeAccess bool   `json:"primeAccess"`
	Stream      bool   `json:"stream"`
	Link        string `json:"link"`
	Update      bool   `json:"update"`
	ID          string `json:"id"`
	AsString    string `json:"asString"`
	Message     string `json:"message"`
	Priority    bool   `json:"priority"`
}

// ToDays : Diff between Today and the given field Date
func (news News) ToDays() string {
	date, err := time.Parse("2006-01-02T15:04:05.000Z", news.Date)
	if err != nil {
		log.Print("Error while parsing date.", err)
		return "0 days ago"
	}
	now := time.Now()

	days := now.Sub(date).Hours() / 24

	return fmt.Sprintf("%d days ago", int64(days))
}
