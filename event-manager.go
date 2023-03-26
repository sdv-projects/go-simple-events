package events

import (
	"context"
	"sync"
)

type eventManager struct {
	handlers map[Event]map[EventHandler]bool
	mu       sync.RWMutex
}

func (ev *eventManager) Subscribe(handler EventHandler, events ...Event) {
	ev.mu.Lock()
	defer ev.mu.Unlock()
	for _, event := range events {
		_, ok_e := ev.handlers[event]
		if !ok_e {
			ev.handlers[event] = make(map[EventHandler]bool)
		}
		_, ok_h := ev.handlers[event][handler]
		if !ok_h {
			ev.handlers[event][handler] = true
		}
	}
}

func (ev *eventManager) Unsubscribe(handler EventHandler, events ...Event) {
	ev.mu.Lock()
	defer ev.mu.Unlock()
	for _, event := range events {
		hm, ok_e := ev.handlers[event]
		if ok_e {
			_, ok_h := hm[handler]
			if ok_h {
				delete(hm, handler)
			}
			if len(hm) == 0 {
				delete(ev.handlers, event)
			}
		}
	}
}

func (ev *eventManager) Invoke(ctx context.Context, event Event, data any) {
	ev.mu.RLock()
	defer ev.mu.RUnlock()
	hm, ok_e := ev.handlers[event]
	if !ok_e {
		return
	}
	for h := range hm {
		go h.OnEvent(ctx, event, data)
	}
}

// The function creates a new EventManager instance.
func NewEventManager() EventManager {
	return &eventManager{
		handlers: make(map[Event]map[EventHandler]bool),
		mu:       sync.RWMutex{},
	}
}
