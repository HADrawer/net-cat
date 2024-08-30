package main

import (
	"chat"
	"fmt"
	"log"
	"net"
	"sync"
)

type Message struct {
	from    string
	payload []byte
}

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	msgch      chan Message
	length     int
	maxClients int
	mu         sync.Mutex
	clients    map[net.Conn]string
	clientMu   sync.Mutex
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		msgch:      make(chan Message, 10),
		length:     0,
		maxClients: 2,
		clients:    make(map[net.Conn]string),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()

	<-s.quitch
	close(s.msgch)

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		s.mu.Lock()
		if s.length >= s.maxClients {
			s.mu.Unlock()
			conn.Write([]byte("The Chat is Maximum"))
			conn.Close()
			fmt.Println("The Chat is Maximum")
			continue
		}
		s.length++
		s.mu.Unlock()

		s.clientMu.Lock()
		s.clients[conn] = ""
		s.clientMu.Unlock()

		fmt.Println("new connection to the server:", conn.RemoteAddr())
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	defer func() {
		s.mu.Lock()
		s.length--
		s.mu.Unlock()

		s.clientMu.Lock()
		delete(s.clients, conn)
		s.clientMu.Unlock()
	}()

	conn.Write(chat.WelcomeMessage())

	nameBuf := make([]byte, 256)
	n, err := conn.Read(nameBuf)
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}

	name := string(nameBuf[:n-1])

	if name == "" {
		conn.Write([]byte("No name provided"))
		return
		} 

		s.clientMu.Lock()
		s.clients[conn] = name
		s.clientMu.Unlock()

		s.broadcast([]byte(name + " joined the Chat \n"))

	buf := make([]byte, 2048)
	for {
		
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("%s left the Chat \n", conn.RemoteAddr())
			s.broadcast([]byte(name +" left the Chat", ))
			break
		}
		
	
		msg := append([]byte("["+name+"]"), buf[:n]...)
		s.broadcast(msg)
	}
}
func(s *Server) broadcast(msg []byte){
	s.clientMu.Lock()
	defer s.clientMu.Unlock()

	for client, name := range s.clients {
	if name != ""{
		client.Write(msg)
	}
	}
}


func main() {
	server := NewServer(":3000")

	// go func() {
	// 	for msg := range server.msgch {
	// 		fmt.Printf("[%s]: %s", msg.from, string(msg.payload))
	// 	}
	// }()
	log.Fatal(server.Start())
}
