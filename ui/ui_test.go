package ui

import (
	"sync"
	"testing"
	"time"

	"github.com/zwh8800/getall/event"
)

func TestBar(t *testing.T) {
	var wg sync.WaitGroup

	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		defer wg.Done()
		event.Server.Publish(event.ProgressStart, 1, "#1")
		for i := 0; i <= 10; i++ {
			event.Server.Publish(event.ProgressUpdate, 1, i*10)
			time.Sleep(300 * time.Millisecond)
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		defer wg.Done()
		event.Server.Publish(event.ProgressStart, 2, "#2")
		for i := 0; i <= 10; i++ {
			event.Server.Publish(event.ProgressUpdate, 2, i*10)
			time.Sleep(200 * time.Millisecond)
		}
	}(&wg)

	time.Sleep(500 * time.Millisecond)

	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		defer wg.Done()
		event.Server.Publish(event.ProgressStart, 3, "#3")
		for i := 0; i <= 10; i++ {
			event.Server.Publish(event.ProgressUpdate, 3, i*10)
			time.Sleep(150 * time.Millisecond)
		}
	}(&wg)

	wg.Wait()
}
