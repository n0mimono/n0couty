package application

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type NotifyHandler struct {
	Forward chan *CrawlState
	Join    chan *Client
	Leave   chan *Client
}

func NewNotifyHandler() *NotifyHandler {
	return &NotifyHandler{
		Forward: make(chan *CrawlState),
		Join:    make(chan *Client),
		Leave:   make(chan *Client),
	}
}

func (hndl *NotifyHandler) Run() {
	clients := map[*Client]bool{}

	for {
		select {
		case client := <-hndl.Join:
			clients[client] = true
		case client := <-hndl.Leave:
			delete(clients, client)
			close(client.Send)
		case msg := <-hndl.Forward:
			data, _ := json.Marshal(msg)
			for client := range clients {
				select {
				case client.Send <- data:
				default:
					delete(clients, client)
					close(client.Send)
				}
			}
		}
	}
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 256,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (hndl *NotifyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	client := &Client{
		Socket: socket,
		Send:   make(chan []byte, 256),
	}

	hndl.Join <- client
	defer func() {
		hndl.Leave <- client
	}()

	client.doWrite()
}

type Client struct {
	Socket *websocket.Conn
	Send   chan []byte
}

func (c *Client) doWrite() {
	for msg := range c.Send {
		if err := c.Socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.Socket.Close()
}
