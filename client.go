package main

type Client struct {
	ChatID int64
}

func NewClient(chatID int64) *Client {
	return &Client{ChatID: chatID}
}

func (client *Client) StartConsultation() {
	if isClientWait[client] {
		sendMessage(client.ChatID, "You are already waiting consultation. Stop current to start a new one!")
		return
	}
	if connectedChatByChatID[client.ChatID] != nil {
		sendMessage(client.ChatID, "Stop current consultation to start a new one!")
		return
	}

	consultant := getConsultantFromWaitList()

	if consultant == nil {
		sendMessage(client.ChatID, "Waiting for free consultants...")
		client.addToWaitList()
		return
	}

	connect(consultant, client)
	consultant.deleteFromWaitList()
	sendMessage(client.ChatID, "Welcome message for clients")
	sendMessage(consultant.ChatID, "Welcome message for consultants")
	return
}

func (client *Client) StopConsultation() {
	consultantChatID := connectedChatByChatID[client.ChatID]
	if consultantChatID == nil {
		sendMessage(client.ChatID, "Can't stop consultation. You haven't started it yet")
		return
	}
	breakConnection(client.ChatID, *consultantChatID)
	sendMessage(client.ChatID, "Goodbye message for client")
	sendMessage(*consultantChatID, "Goodbye message for consultant")
}

func (client *Client) StopWaiting() {
	if !isClientWait[client] {
		sendMessage(client.ChatID, "You are not waiting for consultation now")
		return
	}
	client.deleteFromWaitList()
	sendMessage(client.ChatID, "Goodbye message for client")
}

func (client *Client) addToWaitList() {
	if client != nil {
		isClientWait[client] = true
	}
}

func (client *Client) deleteFromWaitList() {
	delete(isClientWait, client)
}

func getClientFromWaitList() *Client {
	for c := range isClientWait {
		return c
	}
	return nil
}
