package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	go startLoginService()

	bot.Debug = true

	go startSendingMessages(bot)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID

		user := userByChatID[chatID]
		if user == nil {
			user = NewUser(chatID, "client")
		}

		if update.Message.IsCommand() {
			user.HandleCommand(update.Message.Command())
			continue
		}

		connectedChat := connectedChatByChatID[chatID]
		if connectedChat == nil {
			user.HandleCommand("start")
			continue
		}

		sendMessage(*connectedChat, update.Message.Text)
	}

	// todo: shutdown
	// todo: stop messageSender
}

type User interface {
	HandleCommand(command string)
}

func NewUser(chatID int64, role string) User {
	var u User

	switch role {
	case "consultant":
		u = NewConsultant(chatID)
	default:
		u = NewClient(chatID)
	}

	userByChatID[chatID] = u

	return u
}

type message struct {
	ChatID int64
	Text   string
}

func startSendingMessages(bot *tgbotapi.BotAPI) {
	for m := range messages {
		if _, err := bot.Send(tgbotapi.NewMessage(m.ChatID, m.Text)); err != nil {
			log.Println(err)
		}
	}
}

func sendMessage(chatID int64, text string) {
	messages <- message{
		ChatID: chatID,
		Text:   text,
	}
}

// todo: maps concurrent access
var (
	messages = make(chan message)

	isConsultantWait = make(map[*Consultant]bool)
	isClientWait     = make(map[*Client]bool)

	userByChatID = make(map[int64]User)

	connectedChatByChatID = make(map[int64]*int64)
)
