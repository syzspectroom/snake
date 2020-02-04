package transport

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Server struct {
	mux sync.Mutex
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewServer() *Server {
	server := &Server{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}

	go server.run()

	return server
}

func (s *Server) AddClient(conn *websocket.Conn) {
	client := newClient(s, conn)
	fmt.Printf("%+v\n", client)
}

func (s *Server) stopServer() {
	fmt.Println("stop server")
	close(s.broadcast)
	close(s.register)
	close(s.unregister)
}

func (s *Server) removeClient(c *Client) {
	if _, ok := s.clients[c]; ok {
		delete(s.clients, c)
		close(c.send)
	}
}

func (s *Server) Deliver(msg []byte) {
	s.broadcast <- msg

}

func (s *Server) run() {
	defer s.stopServer()
	for {
		select {
		case client := <-s.register:
			fmt.Println("clinet registered")
			s.clients[client] = true
		case client := <-s.unregister:
			fmt.Println("clinet unregistered")
			s.removeClient(client)
		case message := <-s.broadcast:
			fmt.Println("clinet message")
			for client := range s.clients {
				select {
				case client.send <- message:
				default:
					s.removeClient(client)
				}
			}

		}
	}
}
