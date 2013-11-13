package main

// START1 OMIT
type Hub struct {
	// Registered connections. OMIT
	Connections map[*Connection]bool
	// Inbound messages from the connections. OMIT
	Broadcast chan string
	// Register requests from the connections. OMIT
	Register chan *Connection
}
// END1 OMIT

// START2 OMIT
func (h *Hub) Run() {
	for {
		select {
			case c := <-h.Register:
				h.Connections[c] = true
			case m := <-h.Broadcast:
				for c := range h.Connections {
						c.Send <- m
				}
		}
	}
}
// END2 OMIT