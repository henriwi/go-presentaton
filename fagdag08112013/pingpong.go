// +build OMIT

package main

import (
	"fmt"
	"time"
)

// STARTMAIN1 OMIT
type Ball struct{ hits int }

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball)
	time.Sleep(2 * time.Second)
	<-table
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(300 * time.Millisecond)
		table <- ball
	}
}

// STOPMAIN1 OMIT
