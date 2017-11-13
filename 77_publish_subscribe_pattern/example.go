package main

import (
	"errors"
	"time"
)

type Message struct {
	// Contents
}

type Subscription struct {
	ch chan<- Message

	Inbox chan Message
}

func (s *Subscription) Publish(msg Message) error {
	if _, ok := <-s.ch; !ok {
		return errors.New("Topic has been closed")
	}

	s.ch <- msg

	return nil
}

type Topic struct {
	Subscribers    []Session
	MessageHistory []Message
}

func (t *Topic) Subscribe(uid uint64) (Subscription, error) {
	// Get session and create one if it's the first

	// Add session to the Topic & MessageHistory

	// Create a subscription
	return Subscription{}, nil
}

func (t *Topic) Unsubscribe(Subscription) error {
	// Implementation
	return nil
}

func (t *Topic) Delete() error {
	// Implementation
	return nil
}

type User struct {
	ID   uint64
	Name string
}

type Session struct {
	User      User
	Timestamp time.Time
}

func main() {
}
