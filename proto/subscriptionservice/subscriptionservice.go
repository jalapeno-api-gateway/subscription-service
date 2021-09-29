package subscriptionservice

import (
	"context"
	"log"

	"github.com/jalapeno-api-gateway/subscription-service/helpers"
	"github.com/jalapeno-api-gateway/subscription-service/model"
	"github.com/jalapeno-api-gateway/subscription-service/pubsub"
)

type subscriptionServiceServer struct {
	UnimplementedSubscriptionServiceServer
}

func NewServer() *subscriptionServiceServer {
	s := &subscriptionServiceServer{}
	return s
}

func (s *subscriptionServiceServer) SubscribeToLsNodes(subscription *TopologySubscription, responseStream SubscriptionService_SubscribeToLsNodesServer) error {
	log.Printf("SR-App subscribing to LsNodes\n")

	cctx, cancel := context.WithCancel(context.Background())
	sub := pubsub.LsNodeTopic.Subscribe()
	defer func() {
		sub.Unsubscribe()
	}()

	sub.Receive(cctx, func(msg *interface{}) {
		event := (*msg).(model.TopologyEvent)
		if len(subscription.Keys) == 0 || helpers.IsInSlice(subscription.Keys, event.Key) {
			response := convertLsNodeEvent(event)
			err := responseStream.Send(response)
			if err != nil {
				cancel()
			}
		}
	})

	return nil
}

func (s *subscriptionServiceServer) SubscribeToLsLinks(subscription *TopologySubscription, responseStream SubscriptionService_SubscribeToLsLinksServer) error {
	log.Printf("SR-App subscribing to LsLinks\n")

	cctx, cancel := context.WithCancel(context.Background())
	sub := pubsub.LsLinkTopic.Subscribe()
	defer func() {
		sub.Unsubscribe()
	}()

	sub.Receive(cctx, func(msg *interface{}) {
		event := (*msg).(model.TopologyEvent)
		if len(subscription.Keys) == 0 || helpers.IsInSlice(subscription.Keys, event.Key) {
			response := convertLsLinkEvent(event)
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
			event := (*msg).(model.PhysicalInterfaceEvent)
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
		event := (*msg).(model.LoopbackInterfaceEvent)
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