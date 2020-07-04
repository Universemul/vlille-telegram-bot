package main

import (
	//"encoding/json"
	//"fmt"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

var BASE_URL = "https://opendata.lillemetropole.fr/api/records/1.0/search/?dataset=vlille-realtime"

// TODO:
// Infinite loop to get message
func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	var (
		// Universal markup builders.
		menu = &tb.ReplyMarkup{ResizeReplyKeyboard: true}

		// Reply buttons.
		btnLocation = menu.Location("Send location")
	)

	menu.Reply(
		menu.Row(btnLocation),
	)

	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		b.Send(m.Sender, "Please send your location!", menu)
	})
	b.Handle(tb.OnLocation, func(m *tb.Message) {
		res := get_closest_bike(BASE_URL, m.Location.Lat, m.Location.Lng)
		msg := "Liste of stations available within 1km: "
		for _, elem := range res.Result {
			msg = fmt.Sprintf("%s\n\n%s", msg, elem.Display(m.Location.Lat, m.Location.Lng))
		}
		b.Send(m.Sender, msg)
	})
	b.Handle(tb.OnText, func(m *tb.Message) {
		b.Send(m.Sender, "Command not found. Please use /start to get your bike!")
	})
	b.Start()
}
