package telegram

import (
	"context"
	"log"
	"os"

	"github.com/yanzay/tbot/v2"
)

type Bot struct {
	server   *tbot.Server
	client   *tbot.Client
}

func NewBot() Bot {
	bot := tbot.New(os.Getenv("TELEGRAM_BOT_TOKEN"))
	b := &Bot{
		server:   bot,
		client:   bot.Client(),
	}
	bot.HandleMessage("/start", b.helpHandler)
	bot.HandleMessage("/help", b.helpHandler)
	bot.HandleMessage("/update", b.updateHandler)
	bot.HandleMessage("/id", b.IDHandler)
	bot.HandleMessage("/setAuth", b.AuthHandler)
	return *b
}

func (b *Bot) Start(ctx context.Context, stopNotifier chan bool) {
	log.Printf("Bot is going to start")
	err := b.SendMessage("I'm back online!")
	if err != nil {
		log.Printf("Sending online message failed: %v", err)
	}

	errc := make(chan error)
	go func() { errc <- b.server.Start() }()

	select {
	case err := <-errc:
		stopNotifier <- true
		log.Printf("Got an error: %v", err)
	case <-ctx.Done():
		b.SendMessage("Going offline...")
		b.server.Stop()
		log.Printf("Bot went offline")
		<-errc
		stopNotifier <- true
		log.Printf("Context was closed. Reason: %v", ctx.Err())
	}
}