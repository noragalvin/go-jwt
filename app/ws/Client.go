package ws

import (
	"io"
	"log"

	"golang.org/x/net/websocket"
)

const (
	chanBufferSize = 100
)

var (
	clientID = 0
)

// Message ...
type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	UserID  int    `json:"user_id"`
}

// func (self *Message) String() string {
// 	return "> " + self.Body
// }

// Client ...
type Client struct {
	id            int
	conn          *websocket.Conn
	hub           *Hub
	writeChan     chan *Message
	readChan      chan *Message
	doneWriteChan chan bool
	doneReadChan  chan bool
	doneChan      chan bool
}

// NewClient ...
func NewClient(conn *websocket.Conn, hub *Hub) *Client {
	clientID++
	return &Client{
		id:            clientID,
		conn:          conn,
		hub:           hub,
		writeChan:     make(chan *Message, chanBufferSize),
		readChan:      make(chan *Message),
		doneWriteChan: make(chan bool),
		doneReadChan:  make(chan bool),
		doneChan:      make(chan bool),
	}
}

func (c *Client) Write(msg *Message) {
	c.writeChan <- msg
}

func (c *Client) read() {
	var msg Message
	err := websocket.JSON.Receive(c.conn, &msg)
	log.Println("Read msg:", msg)
	if err != nil {
		if err == io.EOF {
			c.doneChan <- true
			return
		}
		c.hub.Err(err)
	} else {
		c.readChan <- &msg
	}
}

// Listen listen read and write chan
func (c *Client) Listen() {
	go c.listenWrite()
	go c.listenRead()
	select {
	case <-c.doneChan:
		c.hub.Del(c)
		c.doneWriteChan <- true
		c.doneReadChan <- true
		return
	}
}

func (c *Client) listenWrite() {
	for {
		select {

		case msg := <-c.writeChan:
			websocket.JSON.Send(c.conn, msg)

		case <-c.doneWriteChan:
			return
		}
	}
}

func (c *Client) listenRead() {
	for {
		go c.read()
		select {

		case msg := <-c.readChan:
			if msg.UserID != 0 {
				c.hub.SendToOne(msg)
				return
			}
			c.hub.SendAll(msg)
		case <-c.doneReadChan:
			return
		}
	}
}
