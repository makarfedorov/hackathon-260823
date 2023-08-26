package main

import (
	"fmt"
)

func (consultant *Consultant) HandleCommand(command string) {
	switch command {
	case "start":
		startCommand(consultant.ChatID)
	case "help":
		messages <- message{
			ChatID: consultant.ChatID,
			Text: "/log_in to log in to the system\n" +
				"/start_consultation start session as consultant\n" +
				"/stop_consultation stop session as consultant\n" +
				"/stop_waiting stop waiting for consultation\n" +
				"/debug debug information",
		}
	case "log_in":
		loginCommand(consultant.ChatID)
	case "debug":
		debugCommand(consultant.ChatID)
	case "start_consultation":
		consultant.StartConsultation()
	case "stop_consultation":
		consultant.StopConsultation()
	case "stop_waiting":
		consultant.StopWaiting()
	default:
		unknownCommand(consultant.ChatID)
	}
}

func (client *Client) HandleCommand(command string) {
	switch command {
	case "start":
		startCommand(client.ChatID)
	case "help":
		messages <- message{
			ChatID: client.ChatID,
			Text: "/log_in to log in to the system\n" +
				"/start_consultation start session as client\n" +
				"/stop_consultation stop session as client\n" +
				"/stop_waiting stop waiting for consultation\n" +
				"/debug debug information",
		}
	case "start_consultation":
		client.StartConsultation()
	case "stop_consultation":
		client.StopConsultation()
	case "stop_waiting":
		client.StopWaiting()
	case "log_in":
		loginCommand(client.ChatID)
	case "debug":
		debugCommand(client.ChatID)
	default:
		unknownCommand(client.ChatID)
	}
}

func startCommand(chatID int64) {
	messages <- message{
		ChatID: chatID,
		Text:   "Type /help to get available commands",
	}
}

func loginCommand(chatID int64) {
	messages <- message{
		ChatID: chatID,
		Text:   fmt.Sprintf("http://www.localhost:8080/login?chat_id=%v", chatID),
	}
}

func unknownCommand(chatID int64) {
	messages <- message{
		ChatID: chatID,
		Text:   "Unknown command",
	}
}

func debugCommand(chatID int64) {
	messages <- message{
		ChatID: chatID,
		Text: fmt.Sprintf("free consultants: %v\nwaiting clients: %v",
			isConsultantWait, isClientWait),
	}
}
