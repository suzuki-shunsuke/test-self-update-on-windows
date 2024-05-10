package main

import (
	"fmt"
	"log"
)

func main() {
	if err := core(); err != nil {
		log.Fatal(err)
	}
}

func core() error {
	fmt.Println("bar")
	return nil
}
