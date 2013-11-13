package main

import (
	"code.google.com/p/go.net/websocket"
	"flag"
	"net/http"
	"text/template"
)

var homeTempl = template.Must(template.ParseFiles("home.html"))
var addr = flag.String("addr", ":8080", "http service address")

func homeHandler(c http.ResponseWriter, req *http.Request) {
	homeTempl.Execute(c, req.Host)
}

// START1 OMIT
var hub = Hub{
	Broadcast:   make(chan string),
	Register:    make(chan *Connection),
	Connections: make(map[*Connection]bool),
}
// END1 OMIT

// START2 OMIT
func main() {
	go hub.Run()
	http.HandleFunc("/", homeHandler)
	http.Handle("/ws", websocket.Handler(wsHandler))
	http.ListenAndServe(*addr, nil)
}
// END2 OMIT