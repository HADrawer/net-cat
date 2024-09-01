package netcat

import (
	"io"
	"log"
	"net"
	"os"
	"sync"
)


type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	length     int
	maxClients int
	mu         sync.Mutex
	clients    map[net.Conn]string
	clientMu   sync.Mutex
	messageHistory []string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		length:     0,
		maxClients: 10,
		clients:    make(map[net.Conn]string),
	}
}

func WelcomeMessage() []byte {
	file, err := os.Open("assets/welcome.txt")
	if err != nil {
		log.Fatal("Error opening txt file")
	}
	defer file.Close()
	result, _ := io.ReadAll(file)
	return result
}
