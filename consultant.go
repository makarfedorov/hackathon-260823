package main

type Consultant struct {
	ChatID int64
}

func NewConsultant(chatID int64) *Consultant {
	return &Consultant{
		chatID,
	}
}

func (consultant *Consultant) StartConsultation() {
	if isConsultantWait[consultant] {
		sendMessage(consultant.ChatID, "You are already waiting consultation. Stop current to start a new one!")
		return
	}
	if connectedChatByChatID[consultant.ChatID] != nil {
		sendMessage(consultant.ChatID, "Stop current consultation to start a new one!")
		return
	}

	client := getClientFromWaitList()

	if client == nil {
		sendMessage(consultant.ChatID, "Waiting for new client...")
		consultant.addToWaitList()
		return
	}

	connect(consultant, client)
	client.deleteFromWaitList()
	sendMessage(consultant.ChatID, "Welcome message for consultants")
	sendMessage(client.ChatID, "Welcome message for clients")
	return
}

func (consultant *Consultant) StopConsultation() {
	clientChatID := connectedChatByChatID[consultant.ChatID]
	if clientChatID == nil {
		sendMessage(consultant.ChatID, "Can't stop consultation. You haven't started it yet")
		return
	}
	breakConnection(consultant.ChatID, *clientChatID)
	sendMessage(consultant.ChatID, "Goodbye message for consultant")
	sendMessage(*clientChatID, "Goodbye message for client")
}

func (consultant *Consultant) StopWaiting() {
	if !isConsultantWait[consultant] {
		sendMessage(consultant.ChatID, "You are not waiting for consultation now")
		return
	}
	consultant.deleteFromWaitList()
	sendMessage(consultant.ChatID, "Goodbye message for consultant")
}

func (consultant *Consultant) addToWaitList() {
	if consultant != nil {
		isConsultantWait[consultant] = true
	}
}

func (consultant *Consultant) deleteFromWaitList() {
	delete(isConsultantWait, consultant)
}

func getConsultantFromWaitList() *Consultant {
	for c := range isConsultantWait {
		return c
	}
	return nil
}
