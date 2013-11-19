package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	_, err := ioutil.ReadFile("does_not_exist.txt")

	if (err != nil) {
		fmt.Println(err)
	}
}