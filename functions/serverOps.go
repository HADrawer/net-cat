package netcat

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	} else {
		fmt.Println("Listening on the port", s.listenAddr)
	}

	defer ln.Close()
	s.ln = ln

	go s.acceptLoop()

	<-s.quitch

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
			conn.Write([]byte("The Chat is Maximum \n"))
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

	conn.Write(WelcomeMessage())
	
	
	nameBuf := make([]byte, 256)
	n, err := conn.Read(nameBuf)
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}

	name := string(nameBuf[:n-1])

	if name == "" {
		conn.Write([]byte("No name provided \n"))
		return
	}
	if s.NameCheck(name) == true {
		conn.Write([]byte("The name is used \n"))
		return
	}
	s.clientMu.Lock()
	for _, msg := range s.messageHistory {
		conn.Write([]byte(msg))
	}
	s.clientMu.Unlock()

	s.clientMu.Lock()
	s.clients[conn] = name
	s.clientMu.Unlock()
	s.broadcast([]byte(name + " joined the Chat \n"))

	buf := make([]byte, 2048)
	for {

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("%s left the Chat \n", conn.RemoteAddr())
			s.broadcast([]byte(name + " left the Chat \n"))
			break
		}
		msg := buf[:n]

		messageStr := string(msg)
		messageStr = strings.TrimSpace(messageStr)
		if len(messageStr) == 0 {
			conn.Write([]byte("The Message is Empty \n"))
			continue
		}
		now := time.Now()
		formatted := now.Format("2006-01-02 15:04:05")
		msg = append([]byte("["+name+"]:"), []byte(messageStr+"\n")...)
		msg = append([]byte("["+formatted +"]") , msg...)
		s.messageHistory = append(s.messageHistory, string(msg))

		s.broadcast(msg)

	}
}
func (s *Server) NameCheck(name string) bool{
	s.clientMu.Lock()
	defer s.clientMu.Unlock()
	for _, name1 := range s.clients {
	if name == name1 {
		return true
	}
	}
	return false
}

func (s *Server) broadcast(msg []byte) {
	s.clientMu.Lock()
	defer s.clientMu.Unlock()

	for client, name := range s.clients {
		if name != "" {
			client.Write(msg)
		}
	}
}
