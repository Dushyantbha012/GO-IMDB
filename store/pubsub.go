package store

import (
	"encoding/json"
	"sync"
)

type MessageType int

const (
	StringMessage MessageType = iota
	JSONMessage
	BinaryMessage
	IntegerMessage
	ArrayMessage
)

type Message struct {
	Type    MessageType
	Content interface{}
}

type PubSub struct {
	mu       sync.Mutex
	channels map[string][]chan Message
}

func NewPubSub() *PubSub {
	return &PubSub{
		channels: make(map[string][]chan Message),
	}
}

func (ps *PubSub) Subscribe(channel string) <-chan Message {
	ch := make(chan Message)
	ps.mu.Lock()
	ps.channels[channel] = append(ps.channels[channel], ch)
	ps.mu.Unlock()
	return ch
}

func (ps *PubSub) PublishString(channel, message string) {
	ps.Publish(channel, Message{Type: StringMessage, Content: message})
}

func (ps *PubSub) PublishJSON(channel string, data interface{}) error {
	// Create a Message with the raw data
	msg := Message{
		Type:    JSONMessage,
		Content: data,
	}
	ps.Publish(channel, msg)
	return nil
}

func (ps *PubSub) PublishBinary(channel string, data []byte) {
	ps.Publish(channel, Message{Type: BinaryMessage, Content: data})
}

func (ps *PubSub) PublishInteger(channel string, number int64) {
	ps.Publish(channel, Message{Type: IntegerMessage, Content: number})
}

func (ps *PubSub) PublishArray(channel string, array []interface{}) {
	ps.Publish(channel, Message{Type: ArrayMessage, Content: array})
}

func (ps *PubSub) Publish(channel string, message Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	// No need to convert message to JSON string here

	for _, ch := range ps.channels[channel] {
		ch <- message
	}
}

// Update MessageToJSON to handle JSON messages properly
func MessageToJSON(msg Message) string {
	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}
