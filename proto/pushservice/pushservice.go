package pushservice

import (
	"context"
	"log"

	"gitlab.ost.ch/ins/jalapeno-api/push-service/helpers"
	"gitlab.ost.ch/ins/jalapeno-api/push-service/model"
	"gitlab.ost.ch/ins/jalapeno-api/push-service/pubsub"
)

type pushServiceServer struct {
	UnimplementedPushServiceServer
}

func NewServer() *pushServiceServer {
	s := &pushServiceServer{}
	return s
}

func (s *pushServiceServer) SubscribeToLsNodes(subscription *LsNodeSubscription, responseStream PushService_SubscribeToLsNodesServer) error {
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
			err := responseStream.Send(&response)
			if err != nil {
				cancel()
			}
		}
	})

	return nil
}

func (s *pushServiceServer) SubscribeToLsLinks(subscription *LsLinkSubscription, responseStream PushService_SubscribeToLsLinksServer) error {
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
			err := responseStream.Send(&response)
			if err != nil {
				cancel()
			}
		}
	})

	return nil
}

func (s *pushServiceServer) SubscribeToTelemetryData(subscription *TelemetrySubscription, responseStream PushService_SubscribeToTelemetryDataServer) error {
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
				response := convertTelemetryEvent(event, subscription.PropertyNames, model.PhysicalInterfaceTelemetryEvent)
				err := responseStream.Send(&response)
				if err != nil {
					cancelA()
				}
			}
		})
	}()

	loopbackInterfaceSubscription.Receive(cctxB, func(msg *interface{}) {
		event := (*msg).(model.LoopbackInterfaceEvent)
		if len(subscription.Ipv4Addresses) == 0 || helpers.IsInSlice(subscription.Ipv4Addresses, event.Ipv4Address) {
			response := convertTelemetryEvent(event, subscription.PropertyNames, model.LoopbackInterfaceTelemetryEvent)
			err := responseStream.Send(&response)
			if err != nil {
				cancelB()
			}
		}
	})
	
	return nil
}

func (s *pushServiceServer) SubscribeToDataRate(subscription *DataRateSubscription, responseStream PushService_SubscribeToDataRateServer) error {
	log.Printf("SR-App subscribing to DataRate\n")

	cctx, cancel := context.WithCancel(context.Background())
	physicalInterfaceSubscription := pubsub.PhysicalInterfaceTopic.Subscribe()

	defer func() {
		physicalInterfaceSubscription.Unsubscribe()
	}()

	physicalInterfaceSubscription.Receive(cctx, func(msg *interface{}) {
		event := (*msg).(model.PhysicalInterfaceEvent)
		if subscription.Ipv4Address == event.Ipv4Address {
			response := DataRate{DataRate: event.DataRate}
			err := responseStream.Send(&response)
			if err != nil {
				cancel()
			}
		}
	})

	return nil
}
