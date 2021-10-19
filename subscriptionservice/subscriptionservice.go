package subscriptionservice

import (
	"context"
	"log"

	"github.com/jalapeno-api-gateway/subscription-service/helpers"
	"github.com/jalapeno-api-gateway/subscription-service/events"
	"github.com/jalapeno-api-gateway/subscription-service/pubsub"
	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
)

type subscriptionServiceServer struct {
	jagw.UnimplementedSubscriptionServiceServer
}

func NewServer() *subscriptionServiceServer {
	s := &subscriptionServiceServer{}
	return s
}

func (s *subscriptionServiceServer) SubscribeToLsNodes(subscription *jagw.TopologySubscription, responseStream jagw.SubscriptionService_SubscribeToLsNodesServer) error {
	log.Printf("SR-App subscribing to LsNodes\n")

	cctx, cancel := context.WithCancel(context.Background())
	sub := pubsub.LsNodeTopic.Subscribe()
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

func (s *subscriptionServiceServer) SubscribeToLsLinks(subscription *jagw.TopologySubscription, responseStream jagw.SubscriptionService_SubscribeToLsLinksServer) error {
	log.Printf("SR-App subscribing to LsLinks\n")

	cctx, cancel := context.WithCancel(context.Background())
	sub := pubsub.LsLinkTopic.Subscribe()
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

func (s *subscriptionServiceServer) SubscribeToLsPrefixes(subscription *jagw.TopologySubscription, responseStream jagw.SubscriptionService_SubscribeToLsPrefixesServer) error {
	log.Printf("SR-App subscribing to LsLinks\n")

	cctx, cancel := context.WithCancel(context.Background())
	sub := pubsub.LsPrefixTopic.Subscribe()
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

func (s *subscriptionServiceServer) SubscribeToLsSrv6Sids(subscription *jagw.TopologySubscription, responseStream jagw.SubscriptionService_SubscribeToLsSrv6SidsServer) error {
	log.Printf("SR-App subscribing to LsLinks\n")

	cctx, cancel := context.WithCancel(context.Background())
	sub := pubsub.LsSrv6SidTopic.Subscribe()
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

func (s *subscriptionServiceServer) SubscribeToTelemetryData(subscription *jagw.TelemetrySubscription, responseStream jagw.SubscriptionService_SubscribeToTelemetryDataServer) error {
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
			if len(subscription.InterfaceIds) == 0 || isSubscribed(subscription.InterfaceIds, event.Hostname, event.LinkID) {
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
		if len(subscription.InterfaceIds) == 0 || isSubscribed(subscription.InterfaceIds, event.Hostname, event.LinkID) {
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

func isSubscribed(interfaceIds []*jagw.InterfaceIdentifier, hostname string, linkId int32) bool {
	for _, interfaceId := range interfaceIds {
		if *interfaceId.Hostname == hostname && *interfaceId.LinkId == linkId {
			return true
		}
	}
	return false
}