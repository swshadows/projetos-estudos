package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("[INFO] Incoming connection: ", ws.RemoteAddr())
	s.conns[ws] = true
	s.readLoop(ws)
}

func (s *Server) handleWSFeed(ws *websocket.Conn) {
	fmt.Println("[INFO] Incoming connection: ", ws.RemoteAddr())

	for {
		payload := fmt.Sprintf("Dados: %d", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep(1 * time.Second)
	}
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("[ERROR] Read error:", err)
			continue
		}
		msg := buf[:n]

		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			_, err := ws.Write(b)
			if err != nil {
				fmt.Println("[ERROR] Write error:", err)
			}
		}(ws)
	}
}

func main() {
	s := NewServer()
	http.Handle("/ws", websocket.Handler(s.handleWS))
	http.Handle("/feed", websocket.Handler(s.handleWSFeed))
	http.ListenAndServe(":3000", nil)
}
