# go-simple-events

**go-simple-events** is a simple and lightweight implementation of the event system in Golang. The implementation does not guarantee the order in which events are handled.

# Usage example
You need to do the following actions: 
- Define event constants and event data types:
```go
// Event definitions
const (
	MY_EVENT1 events.Event = "MY_EVENT1"
	MY_EVENT2 events.Event = "MY_EVENT2"
)

// Data type for MY_EVENT1
type MyEvent1Data struct {
	Data int32
}

// Data type for MY_EVENT2
type MyEvent2Data struct {
	Time time.Time
}
```
- Implement *EventHandler* interface:
```go
type MyEventHandler struct {
	// some data ...
}

// Implementing an event handling method
func (h *MyEventHandler) OnEvent(ctx context.Context, event events.Event, data any) {
	switch event {
	case MY_EVENT1:
		if event1-data := data.(*MyEvent1Data); event1-data != nil {
			// ...
		}

	case MY_EVENT2:
		if event2-data := data.(*MyEvent2Data); event2-data != nil {
			// ...
		}
	}
}
```
- Subscribe your EventHandler instance to events:
```go
handler := MyEventHandler{/*...*/}
events.Subscribe(&handler, MY_EVENT1, MY_EVENT2)
```
- Invoke your events:
```go
events.Invoke(ctx, MY_EVENT1, &MyEvent1Data{Data: 15})
events.Invoke(ctx, MY_EVENT2, &MyEvent2Data{Time: time.Now()})
```

If you want to exclude events from handling, you can call the *Unsubscribe* method:
```go
events.Unsubscribe(&handler, MY_EVENT1)
```
