package main

import (
	"fmt"
	"github.com/tixu/coffe-shop/barista"
	"github.com/tixu/coffe-shop/customer"
	"github.com/tixu/coffe-shop/store"
	"time"
)

// Purely for CLI purposes, so func main has something to start
var startCLI = func() {
	outputChan := make(chan string)
	s := store.New()
	s.CloseAfter(10 * time.Second)
	s.Customers(customer.RandomGroupOf(2))
	s.Baristas(barista.RandomGroupOf(2))
	go func() {
		err := s.Open(outputChan)
		if err != nil {
			panic(err)
		} else {
			outputChan <- "*** Store is closed now ***"

		}
	}()
	for {
		select {
		case str := <-outputChan:
			fmt.Println(str)
		case <-time.After(10 * time.Second):
			return
		}
	}
}

func main() {
	startCLI()
}
