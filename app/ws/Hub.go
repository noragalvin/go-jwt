package ws

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// Hub ...
type Hub struct {
	messages   []*Message
	clients    map[int]*Client
	register   chan *Client
	unregister chan *Client
	sendToAll  chan *Message
	sendToOne  chan *Message
	doneCh     chan bool
	errCh      chan error
}

// NewHub create new object Hub
func NewHub() *Hub {
	return &Hub{
		messages:   make([]*Message, 0),
		clients:    make(map[int]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		sendToAll:  make(chan *Message),
		sendToOne:  make(chan *Message),
		doneCh:     make(chan bool),
		errCh:      make(chan error),
	}
}

// Add new client to channel
func (s *Hub) Add(c *Client) {
	s.register <- c
	//TODO:store to database
}

// Del a client from channel
func (s *Hub) Del(c *Client) {
	s.unregister <- c

}

// SendAll : send to every clients
func (s *Hub) SendAll(msg *Message) {
	s.sendToAll <- msg
}

// SendToOne : send to specific client
func (s *Hub) SendToOne(msg *Message) {
	s.sendToOne <- msg
}

// Err show err
func (s *Hub) Err(err error) {
	s.errCh <- err
}

func (s *Hub) sendPastMessages(c *Client) {
	go func() {
		for _, msg := range s.messages {
			c.Write(msg)
		}
	}()
}

func (s *Hub) sendAll(msg *Message) {
	for _, c := range s.clients {
		go c.Write(msg)
	}
}

func (s *Hub) sendToAnotherUser(msg *Message) {
	for _, c := range s.clients {
		if c.id == msg.UserID {
			go c.Write(msg)
		}
	}
}

// Handler control the websocket
func (s *Hub) Handler() http.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {
		defer func() {
			if err := ws.Close(); err != nil {
				s.errCh <- err
			}
		}()

		client := NewClient(ws, s)
		s.Add(client)
		client.Listen()
	})
}

// Run websocket
func (s *Hub) Run() {
	for {
		select {

		case c := <-s.register:
			s.clients[c.id] = c
			log.Println("Added new client. Clients connected:", len(s.clients))
			s.sendPastMessages(c)

		case c := <-s.unregister:
			log.Println("Client:", c.id, " is quited")
			delete(s.clients, c.id)

		case msg := <-s.sendToAll:
			s.messages = append(s.messages, msg)
			s.sendAll(msg)

		case msg := <-s.sendToOne:
			log.Println("Send to one")
			s.sendToAnotherUser(msg)

		case err := <-s.errCh:
			log.Println("Error:", err.Error())

		case <-s.doneCh:
			return
		}
	}
}
