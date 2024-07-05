package main

import (
	"fmt"
	"log"
	"example.com/greeting"

)

func main(){

	log.SetPrefix("greetings:")
	log.SetFlags(0)

	names := []string{"bhuws", "sidu", "vedu"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
