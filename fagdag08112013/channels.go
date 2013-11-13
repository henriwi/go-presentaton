package main

import "fmt"

func main() {
	//START1 OMIT
	//Oppretter kanal
	channel := make(chan int)

	// Sender melding på en kanal
	go channel <- 1

	// Venter på svar og assigner til en variabel
	svar := <- channel
	//END1 OMIT
}