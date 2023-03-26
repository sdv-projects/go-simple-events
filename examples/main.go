package main

import (
	"context"
	"log"
	"sync"
	"time"

	events "github.com/sdv-projects/go-simple-events"
)

type contextKey string

const (
	MY_EVENT1 events.Event = "MY_EVENT1"
	MY_EVENT2 events.Event = "MY_EVENT2"
	WG_KEY    contextKey   = "wg"
)

type EventData struct {
	Value int
}

type MyEventHandler struct{}

func (h *MyEventHandler) OnEvent(ctx context.Context, event events.Event, data any) {
	log.Printf("[OnEvent] Event: %s, Data: %v, In process...", event, data)
	time.Sleep(5 * time.Second)
	log.Printf("[OnEvent] Event: %s, Data: %v, DONE!", event, data)

	wg := ctx.Value(WG_KEY).(*sync.WaitGroup)
	wg.Done()
}

func main() {
	ctx := context.Background()
	var wg sync.WaitGroup

	handler := MyEventHandler{}
	events.Subscribe(&handler, MY_EVENT1, MY_EVENT2)

	log.Println("Invoke event: MY_EVENT1")
	wg.Add(1)
	events.Invoke(context.WithValue(ctx, WG_KEY, &wg), MY_EVENT1, &EventData{Value: 1})

	log.Println("Invoke event: MY_EVENT2")
	wg.Add(1)
	events.Invoke(context.WithValue(ctx, WG_KEY, &wg), MY_EVENT2, &EventData{Value: 2})

	log.Println("Waiting for events to be handled...")
	wg.Wait()
	log.Println("Events are handled")
}
