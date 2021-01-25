package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {

	s := NewServer("localhost", "8081", newKeycloak())
	s.listen()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

}