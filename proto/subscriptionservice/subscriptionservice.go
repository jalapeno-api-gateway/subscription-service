package subscriptionservice

import (
	"context"
	"log"

	"github.com/jalapeno-api-gateway/subscription-service/helpers"
	"github.com/jalapeno-api-gateway/subscription-service/events"
	"github.com/jalapeno-api-gateway/subscription-service/pubsub"
)

type subscriptionServiceServer struct {
	UnimplementedSubscriptionServiceServer
}

func NewServer() *subscriptionServiceServer {
	s := &subscriptionServiceServer{}
	return s
}

func (s *subscriptionServiceServer) SubscribeToLSNodes(subscription *TopologySubscription, responseStream SubscriptionService_SubscribeToLSNodesServer) error {
	log.Printf("SR-App subscribing to LsNodes\n")

	cctx, cancel := context.WithCancel(context.Background())
	sub := pubsub.LSNodeTopic.Subscribe()
	defer func() {
		sub.Unsubscribe()
	}()

	sub.Receive(cctx, func(msg *interface{}) {
		event := (*msg).(events.TopologyEvent)
		if len(subscription.Keys) == 0 || helpers.IsInSlice(subscription.Keys, event.Key) {
			response := convertLSNodeEvent(event)
			err := responseStream.Send(response)
			if err != nil {
				cancel()
			}
		}
	})

	return nil
}

func (s *subscriptionServiceServer) SubscribeToLSLinks(subscription *TopologySubscription, responseStream SubscriptionService_SubscribeToLSLinksServer) error {
	log.Printf("SR-App subscribing to LsLinks\n")

	cctx, cancel := context.WithCancel(context.Background())
	sub := pubsub.LSLinkTopic.Subscribe()
	defer func() {
		sub.Unsubscribe()
	}()

	sub.Receive(cctx, func(msg *interface{}) {
		event := (*msg).(events.TopologyEvent)
		if len(subscription.Keys) == 0 || helpers.IsInSlice(subscription.Keys, event.Key) {
			response := convertLSLinkEvent(event)
			err := responseStream.Send(response)
			if err != nil {
				cancel()
			}
		}
	})

	return nil
}

func (s *subscriptionServiceServer) SubscribeToLSPrefixes(subscription *TopologySubscription, responseStream SubscriptionService_SubscribeToLSPrefixesServer) error {
	log.Printf("SR-App subscribing to LsLinks\n")

	cctx, cancel := context.WithCancel(context.Background())
	sub := pubsub.LSPrefixTopic.Subscribe()
	defer func() {
		sub.Unsubscribe()
	}()

	sub.Receive(cctx, func(msg *interface{}) {
		event := (*msg).(events.TopologyEvent)
		if len(subscription.Keys) == 0 || helpers.IsInSlice(subscription.Keys, event.Key) {
			response := convertLSPrefixEvent(event)
			err := responseStream.Send(response)
			if err != nil {
				cancel()
			}
		}
	})

	return nil
}

func (s *subscriptionServiceServer) SubscribeToLSSRv6SIDs(subscription *TopologySubscription, responseStream SubscriptionService_SubscribeToLSSRv6SIDsServer) error {
	log.Printf("SR-App subscribing to LsLinks\n")

	cctx, cancel := context.WithCancel(context.Background())
	sub := pubsub.LSSRv6SIDTopic.Subscribe()
	defer func() {
		sub.Unsubscribe()
	}()

	sub.Receive(cctx, func(msg *interface{}) {
		event := (*msg).(events.TopologyEvent)
		if len(subscription.Keys) == 0 || helpers.IsInSlice(subscription.Keys, event.Key) {
			response := convertLSSRv6SIDEvent(event)
			err := responseStream.Send(response)
			if err != nil {
				cancel()
			}
		}
	})

	return nil
}

func (s *subscriptionServiceServer) SubscribeToTelemetryData(subscription *TelemetrySubscription, responseStream SubscriptionService_SubscribeToTelemetryDataServer) error {
	log.Printf("SR-App subscribing to TelemetryData\n")

	cctxA, cancelA := context.WithCancel(context.Background())
	cctxB, cancelB := context.WithCancel(context.Background())
	physicalInterfaceSubscription := pubsub.PhysicalInterfaceTopic.Subscribe()
	loopbackInterfaceSubscription := pubsub.LoopbackInterfaceTopic.Subscribe()
	
	defer func() {
		physicalInterfaceSubscription.Unsubscribe()
		loopbackInterfaceSubscription.Unsubscribe()
	}()

	go func() {
		physicalInterfaceSubscription.Receive(cctxA, func(msg *interface{}) {
			event := (*msg).(events.PhysicalInterfaceEvent)
			if len(subscription.Ipv4Addresses) == 0 || helpers.IsInSlice(subscription.Ipv4Addresses, event.Ipv4Address) {
				response := convertPhysicalInterfaceEvent(event, subscription.PropertyNames)
				err := responseStream.Send(response)
				if err != nil {
					cancelA()
				}
			}
		})
	}()

	loopbackInterfaceSubscription.Receive(cctxB, func(msg *interface{}) {
		event := (*msg).(events.LoopbackInterfaceEvent)
		if len(subscription.Ipv4Addresses) == 0 || helpers.IsInSlice(subscription.Ipv4Addresses, event.Ipv4Address) {
			response := convertLoopbackInterfaceEvent(event, subscription.PropertyNames)
			err := responseStream.Send(response)
			if err != nil {
				cancelB()
			}
		}
	})
	// TODO return status.Errorf(codes.)
	return nil
}