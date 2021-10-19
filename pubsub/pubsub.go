package pubsub

import (
	"context"
	"log"
	"sync"

	"github.com/google/uuid"
)

type Subscription struct {
	id string
	events chan *interface{}
	abort chan struct{}
	topic *threadSafeTopic
}

type threadSafeTopic struct {
	sync.Mutex
	subscriptions []*Subscription
}

var LsNodeTopic *threadSafeTopic
var LsLinkTopic *threadSafeTopic
var LsPrefixTopic *threadSafeTopic
var LsSrv6SidTopic *threadSafeTopic
var PhysicalInterfaceTopic *threadSafeTopic
var LoopbackInterfaceTopic *threadSafeTopic

func InitializeTopics() {
	LsNodeTopic = &threadSafeTopic{}
	LsLinkTopic = &threadSafeTopic{}
	LsPrefixTopic = &threadSafeTopic{}
	LsSrv6SidTopic = &threadSafeTopic{}
	PhysicalInterfaceTopic = &threadSafeTopic{}
	LoopbackInterfaceTopic = &threadSafeTopic{}
}

func (topic *threadSafeTopic) Subscribe() *Subscription {
	s := topic.createNewSubscription()

	topic.Lock()
	defer topic.Unlock()

	topic.subscriptions = append(topic.subscriptions, s)
	return s
}

func (s *Subscription) Unsubscribe() {
	s.topic.Lock()
	defer s.topic.Unlock()

	for index, subscription := range s.topic.subscriptions { // Find subscription index in array
		if subscription.id == s.id {
			s.topic.subscriptions = append(s.topic.subscriptions[:index], s.topic.subscriptions[index+1:]...) // Remove subscription from array
			return
		}
	}
	log.Fatalf("Error when trying to remove subscription: %s from slice", s.id)
}

func (topic *threadSafeTopic) createNewSubscription() *Subscription {
	return &Subscription{
		id: uuid.New().String(),
		events: nil,
		abort: nil,
		topic: topic,
	}
}

func (topic *threadSafeTopic) Publish(event interface{}) {
	topic.Lock()
	defer topic.Unlock()

	for _, subscription := range topic.subscriptions {
		select {
			case <-subscription.abort:
			case subscription.events <-&event:
		}
	}
}

func (subscription *Subscription) Receive(ctx context.Context, callback func(event *interface{})) {
	initializeSubscription(subscription)
	
	loop:
	for {
		select {
			case <-ctx.Done():
				close(subscription.abort)
				break loop
			case event := <-subscription.events:
				callback(event)
		}
	}
}

func initializeSubscription(subscription *Subscription) {
	subscription.events = make(chan *interface{})
	subscription.abort = make(chan struct{})
}