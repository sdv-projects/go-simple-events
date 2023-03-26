package events

import "context"

var defaultEventMgr = NewEventManager()

// The function associates a handler with events.
func Subscribe(handler EventHandler, events ...Event) {
	defaultEventMgr.Subscribe(handler, events...)
}

// The function removes associations between handler and events.
func Unsubscribe(handler EventHandler, events ...Event) {
	defaultEventMgr.Unsubscribe(handler, events...)
}

// The function invokes an event with data.
func Invoke(ctx context.Context, event Event, data any) {
	defaultEventMgr.Invoke(ctx, event, data)
}
