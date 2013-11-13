package main

import (
	"code.google.com/p/go.net/websocket"
)

// START1 OMIT
type Connection struct {
	// The websocket connection. OMIT
	WebSocket *websocket.Conn
	// Buffered channel of outbound messages. OMIT
	Send chan string
}
// END1 OMIT

// START2 OMIT
func (connecion *Connection) Reader() {
	for {
		var message string
		err := websocket.Message.Receive(connecion.WebSocket, &message)
		if err != nil {
			break
		}
		hub.Broadcast <- message
	}
	connecion.WebSocket.Close()
}
// END2 OMIT

// START3 OMIT
func (connection *Connection) Writer() {
	for message := range connection.Send {
		err := websocket.Message.Send(connection.WebSocket, message)
		if err != nil {
			break
		}
	}
	connection.WebSocket.Close()
}
// END3 OMIT

// START4 OMIT
func wsHandler(ws *websocket.Conn) {
	connection := &Connection{Send: make(chan string, 256), WebSocket: ws}
	hub.Register <- connection
	go connection.Writer()
	connection.Reader()
}
// END4 OMIT