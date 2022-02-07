package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Adapter struct {
	bot     *tgbotapi.BotAPI
	ChanMsg chan *tgbotapi.Message
}

func NewTgAdapter(token string) (*Adapter, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	result := Adapter{
		bot:     bot,
		ChanMsg: make(chan *tgbotapi.Message),
	}
	//bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return &result, nil
}

func (a *Adapter) Listen() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := a.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message != nil { // If we got a message
			a.ChanMsg <- update.Message
		}
	}
	return nil
}

func (a *Adapter) Send(msg tgbotapi.Chattable) error {
	_, err := a.bot.Send(msg)
	return err
}
