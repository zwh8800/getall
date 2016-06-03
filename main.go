package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	evbus "github.com/asaskevich/EventBus"

	"github.com/zwh8800/getall/event"
)

func init() {
	evbus.Subscribe(event.Finish, finish)
}

func finish() {
	os.Exit(0)
}

func main() {
	startServices()
	log.Println("services started")

	handleSignal()
	log.Println("signal received")

	stopServices()
	log.Println("gracefully shutdown")
}

func startServices() {

}

func handleSignal() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Kill, os.Interrupt, syscall.SIGTERM)
	<-signalChan
}

func stopServices() {

}
