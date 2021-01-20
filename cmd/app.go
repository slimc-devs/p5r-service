package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var sig chan os.Signal = make(chan os.Signal)

func main() {
	log.Println("Hello world!")
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
