package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"./warframe"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func boolToRedeable(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}

func main() {
	ordisQuotes := [16]string{
		"Operator? Ordis wonders... what are you thinking about?",
		"I've been thinking, Operator...I thought you'd want to know.",
		"Operator, I hope you are comfortable? No...we do not seem to have any seats.",
		"Everything in Ordis, Operator? Is that a pun?! Hmm.... I will attempt to bypass this fault.",
		"Ordis has been counting stars, Operator. All accounted for.",
		"Operator, I will never betray you. I will keep the Orbiter hidden in the void. You can count on me!",
		"Operator, I've been thinking. My misplaced memories and damaged communication systems. What if...Ordis did those things?",
		"Operator, were you visualizing a bloody battle? -Me too!",
		"Ordis is hap - angry. Hmm, I may require maintenance after all.",
		"Operator, the system needs you. Will you begin another mission?",
		"Operator! Did you hear that? It said-- Cosmic background radiation is a riot!",
		"Stand by while I analyze the intelligence profile of the Grineer. Error, not a number! Did the Operator enjoy this witticism?",
		"Do you remember the Old War, Operator? Ordis seems to have... misplaced those memories.",
		"Do not lift the veil. Do not show the door. Do not split the dream.",
		"Maintain the habitat. Maintain the Operator. Mobilize the Tenno.",
		"You are the Tenno. You are the Operator. Ordis is the Cephalon. Ordis is the ship.",
	}

	bot, err := tgbotapi.NewBotAPI("***REMOVED***")
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, errTwo := bot.GetUpdatesChan(u)
	if errTwo != nil {
		panic(errTwo)
	}

	api := warframe.CreateAPI()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		rand.Seed(time.Now().UnixNano())

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.ParseMode = "HTML"

		switch update.Message.Command() {
		case "start":
			msg.Text = ordisQuotes[rand.Intn(len(ordisQuotes))]
		case "quotes":
			msg.Text = ordisQuotes[rand.Intn(len(ordisQuotes))]
		case "latest":
			news := api.GetNews()
			latest := len(news) - 1
			msg.Text = fmt.Sprintf("<b>%s:</b> %s\n<a href=\"%s\">Link</a>\n", news[latest].ToDays(), news[latest].Message, news[latest].Link)
		case "news":
			news := api.GetNews()
			var text string
			text += "<b>News </b>\n"
			for i := 0; i < len(news); i++ {
				text += fmt.Sprintf("<b>%s:</b> %s\n<a href=\"%s\">Link</a>\n", news[i].ToDays(), news[i].Message, news[i].Link)
			}
			msg.Text = text
		case "vallis":
			vallis := api.GetVallisCycle()
			msg.Text = fmt.Sprintf("<b>Vallis (Fortuna)</b>\nCurrent state:\t%s\nTime left:\t<b>%s</b>", vallis.GetCurrentState(), vallis.TimeLeft)
		case "cetus":
			cetus := api.GetCetusCycle()
			msg.Text = fmt.Sprintf("<b>Cetus (Plains of Eidolon)</b>\nCurrent state:\t%s\nTime left:\t<b>%s</b>", cetus.GetCurrentState(), cetus.TimeLeft)
		// case "alerts":
		// case "nightwave":
		// case "fissures":
		default:
			msg.Text = ordisQuotes[0]
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
