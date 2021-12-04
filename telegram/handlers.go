package telegram

import (
	"Corona_Test/network"
	"Corona_Test/test"
	"fmt"
	"log"
	"os"

	"github.com/yanzay/tbot/v2"
)

//// Send this to BotFather to set commands (with /setcommands)
//help - Get this help message
//update - Get new updates
//info - Get info about all tests
//id - Get your own message ID
//setAuth - Set the auth token

// Handle the /start and /help commands here
func (b *Bot) helpHandler(m *tbot.Message) {
	log.Printf("user %s %s with id: %s sent message: '%s'", m.Chat.FirstName, m.Chat.LastName, m.Chat.ID, m.Text)
	msg := `This is a bot whose purpose is to send updates about a corona test.
Commands:
1. Use /help to get this help message
2. Use /update to get an update.
3. Use /info to get info about all tests.
4. Use /id to get your message ID.
5. Use /setAuth to set the auth token`

	if _, err := b.client.SendMessage(m.Chat.ID, msg); err != nil {
		log.Printf("failed to send help message to client: %s", err)
	}
}

// IDHandler Handle the /getID command here
func (b *Bot) IDHandler(m *tbot.Message) {
	if _, err := b.client.SendMessage(m.Chat.ID, "Your ID is: " + m.Chat.ID); err != nil {
		log.Printf("failed to send id message to client: %s", err)
	}
}

// AuthHandler Handle the /setAuth command here
func (b *Bot) AuthHandler(m *tbot.Message) {
	if m.Chat.ID == os.Getenv("TELEGRAM_CHAT_ID") {
		os.Setenv("BEARER", m.Text[9:])
		log.Printf("Auth token was updated")
		if _, err := b.client.SendMessage(m.Chat.ID, "New auth token is: " + os.Getenv("BEARER")); err != nil {
			log.Printf("failed to send auth message to client: %s", err)
		}
	} else {
		if _, err := b.client.SendMessage(m.Chat.ID, "I don't know you, so you don't have this permission"); err != nil {
			log.Printf("failed to send auth without permission message to client: %s", err)
		}
	}
}

func (b *Bot) SendMessage(message string) error {
	if _, err := b.client.SendMessage(
		os.Getenv("TELEGRAM_CHAT_ID"),
		message,
	); err != nil {
		return fmt.Errorf("failed to send normal message to client: %w", err)
	}
	return nil
}

func (b *Bot) updateHandler(m *tbot.Message) {
	err := b.SendMessage("Running...")
	if err != nil {
		log.Printf("Sending the running message went wrong: %s\n", err)
	}

	result, err := network.GetUpdate()
	if err != nil {
		errString := fmt.Errorf("getting the update went wrong: %w", err)
		log.Print(errString)

		b.SendMessage(errString.Error())
		return
	}
	message := test.GetInterestingMessage(result)

	if message != "" {
		err := b.SendMessage(message)
		if err != nil {
			log.Printf("Sending the message went wrong: %s\n", err)
			return
		}
	}

	err = b.SendMessage("Done")
	if err != nil {
		log.Printf("Sending the done message went wrong: %s\n", err)
		return
	}
}