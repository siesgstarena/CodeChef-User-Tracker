package helper

import "github.com/gorilla/websocket"

func SendMessage(conn *websocket.Conn, message string) {
	conn.WriteMessage(websocket.TextMessage, []byte(message))
}
