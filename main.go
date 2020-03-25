package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"./warframe"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func loadQuotes() []string {
	quotes, err := ioutil.ReadFile("quotes.json")
	if err != nil {
		panic(err)
	}
	var result []string
	json.Unmarshal(quotes, &result)
	return result
}

func main() {

	ordisQuotes := loadQuotes()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
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
		case "quote":
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
		case "nightwave":
			nw := api.GetNightwaveInfo()
			var text string
			text += fmt.Sprintf("<b>Nightwave Challenges Season %d</b>\n\n", nw.Season)
			for i := 0; i < len(nw.ActiveChallenges); i++ {
				text += fmt.Sprintf("<b>%d</b> [%s]\n<i>%s</i>\n\n", nw.ActiveChallenges[i].Reputation, nw.ActiveChallenges[i].Title, nw.ActiveChallenges[i].Desc)
			}
			msg.Text = text
		// case "fissures":
		default:
			msg.Text = ordisQuotes[0]
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
