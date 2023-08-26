package main

func connect(consultant *Consultant, client *Client) {
	connectedChatByChatID[consultant.ChatID] = &client.ChatID
	connectedChatByChatID[client.ChatID] = &consultant.ChatID
}

func breakConnection(consultantChatID, clientChatID int64) {
	delete(connectedChatByChatID, consultantChatID)
	delete(connectedChatByChatID, clientChatID)
}
