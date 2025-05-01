package sse

import "painellembretes/models"

type SSEClient chan models.Reminder

type SSEHub struct {
	clients   map[SSEClient]bool
	add       chan SSEClient
	remove    chan SSEClient
	Broadcast chan models.Reminder
}

func NewSSEHub() *SSEHub {
	return &SSEHub{
		clients:   make(map[SSEClient]bool),
		add:       make(chan SSEClient),
		remove:    make(chan SSEClient),
		Broadcast: make(chan models.Reminder),
	}
}

func (hub *SSEHub) Run() {
	for {
		select {
		case client := <-hub.add:
			hub.clients[client] = true
		case client := <-hub.remove:
			delete(hub.clients, client)
			close(client)
		case msg := <-hub.Broadcast:
			for client := range hub.clients {
				select {
				case client <- msg:
				default:
					delete(hub.clients, client)
					close(client)
				}
			}
		}
	}
}
