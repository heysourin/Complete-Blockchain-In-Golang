package main

import (
	"fmt"
	"log" //shows Date()
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	a := "test"
	fmt.Println(a)

	log.Println("test log")
}
