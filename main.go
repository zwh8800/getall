package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/zwh8800/getall/event"
	"github.com/zwh8800/getall/spider"
	_ "github.com/zwh8800/getall/ui"
)

func init() {
	event.Server.Subscribe(event.Finish, finish)
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
	spider.Go()
}

func handleSignal() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Kill, os.Interrupt, syscall.SIGTERM)
	<-signalChan
}

func stopServices() {
	spider.Stop()
}
