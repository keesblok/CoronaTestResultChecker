package main

import (
	"Corona_Test/network"
	"Corona_Test/telegram"
	"Corona_Test/test"
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	bot := telegram.NewBot()

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	stopNotifier := make(chan bool)
	stopServer := make(chan bool)

	go func() {
		for {
			select {
			case <- stopNotifier:
				return
			case <- ticker.C:
				result := network.GetUpdate()
				message := test.GetInterestingMessage(result)

				if message != "" {
					err := bot.SendMessage(message)
					if err != nil {
						log.Printf("Sending the message went wrong: %s\n", err)
						return
					}
				}
			}
		}
	}()

	bot.Start(ctx, stopNotifier, stopServer)

	network.StartServer(stopServer)
}

