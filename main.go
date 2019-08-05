package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/henglory/line-webhook/server"
)

func main() {
	s := server.NewServer()
	s.Start()
	defer s.Close()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Server exiting")
}
