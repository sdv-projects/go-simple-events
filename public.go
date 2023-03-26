package events

import "context"

// The type defines events.
// Add event constants of Event type to your project.
type Event string

// Event manager interface.
// If you need to create a separate EventManager instances, you must use the NewEventManager function.
type EventManager interface {
	// The function associates a handler with events.
	Subscribe(handler EventHandler, events ...Event)
	// The function removes associations between handler and events.
	Unsubscribe(handler EventHandler, events ...Event)
	// The function invokes an event with data.
	Invoke(ctx context.Context, event Event, data any)
}

// Event handler interface.
type EventHandler interface {
	// Event handling function. It's called when an event occurs.
	OnEvent(ctx context.Context, event Event, data any)
}
