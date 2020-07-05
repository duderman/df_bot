package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

type channel string

func (c channel) Recipient() string {
	return string(c)
}

func main() {
	var comment string
	var b *tb.Bot
	var err error

	b, err = tb.NewBot(tb.Settings{
		Token:   os.Getenv("TELEGRAM_API_TOKEN"),
		Poller:  &tb.LongPoller{Timeout: 15 * time.Second},
		Verbose: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	b.Handle(tb.OnText, func(m *tb.Message) {
		comment = m.Text
	})

	b.Handle(tb.OnAnimation, func(m *tb.Message) {
		m.Animation.Caption = comment
		send(b, m.Animation)
		comment = ""
	})

	b.Handle(tb.OnPhoto, func(m *tb.Message) {
		m.Photo.Caption = comment
		send(b, m.Photo)
		comment = ""
	})

	b.Handle(tb.OnVideo, func(m *tb.Message) {
		m.Video.Caption = comment
		send(b, m.Video)
		comment = ""
	})

	fmt.Println("Starting...")
	b.Start()
}

func send(b *tb.Bot, attach interface{}) {
	sendTo := channel(os.Getenv("CHANNEL_NAME"))
	_, err := b.Send(sendTo, attach)
	if err != nil {
		log.Println(err)
	}
}
