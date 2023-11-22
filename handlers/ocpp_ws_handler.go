package handlers

import (
	"fmt"
	"github.com/aliakseizhurauliou/OCPPGolang/mapping"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	Subprotocols: []string{"ocpp1.6"},
	CheckOrigin: func(r *http.Request) bool {
		return true // Пропускаем любой запрос
	},
}

type Server struct {
	clients map[*websocket.Conn]bool
	mu      *sync.Mutex
}

func StartServer() *Server {
	server := Server{
		make(map[*websocket.Conn]bool),
		&sync.Mutex{},
	}

	http.HandleFunc("/", server.ReadMessage)
	go func() {
		http.ListenAndServe(":8080", nil) // Уводим http сервер в горутину
	}()

	return &server
}

func (server *Server) ReadMessage(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)
	defer connection.Close()

	server.mu.Lock()
	server.clients[connection] = true
	server.mu.Unlock()                       // Сохраняем соединение, используя его как ключ
	defer delete(server.clients, connection) // Удаляем соединение

	for {
		mt, message, err := connection.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			break // Выходим из цикла, если клиент пытается закрыть соединение или связь прервана
		}

		ocpp := mapping.Map(string(message))

		if ocpp.MessageType == "bootNotification" {
			fmt.Println("BOOT!")
		}

	}
}

func (server *Server) WriteMessage(message []byte) {
	for conn := range server.clients {
		conn.WriteMessage(websocket.TextMessage, message)
	}
}
