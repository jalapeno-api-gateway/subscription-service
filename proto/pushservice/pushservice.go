package pushservice

import (
	"log"

	"gitlab.ost.ch/ins/jalapeno-api/push-service/helpers"
	"gitlab.ost.ch/ins/jalapeno-api/push-service/subscribers"
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

	events := make(chan subscribers.LsNodeEvent)
	subscribers.SubscribeToLsNodeEvents(events)
	defer func() {
		subscribers.UnSubscribeFromLsNodeEvents(events)
	}()

	for {
		event := <-events
		if len(subscription.Keys) == 0 || helpers.IsInSlice(subscription.Keys, event.Key) {
			response := convertToGrpcLsNodeEvent(event)
			err := responseStream.Send(&response)
			if err != nil {
				return err
			}
		}
	}
}

func (s *pushServiceServer) SubscribeToLsLinks(subscription *LsLinkSubscription, responseStream PushService_SubscribeToLsLinksServer) error {
	log.Printf("SR-App subscribing to LsLinks\n")

	events := make(chan subscribers.LsLinkEvent)
	subscribers.SubscribeToLsLinkEvents(events)
	defer func() {
		subscribers.UnSubscribeFromLsLinkEvents(events)
	}()

	for {
		event := <-events
		if len(subscription.Keys) == 0 || helpers.IsInSlice(subscription.Keys, event.Key) {
			response := convertToGrpcLsLinkEvent(event)
			err := responseStream.Send(&response)
			if err != nil {
				return err
			}
		}
	}
}

func (s *pushServiceServer) SubscribeToTelemetryData(subscription *TelemetrySubscription, responseStream PushService_SubscribeToTelemetryDataServer) error {
	log.Printf("SR-App subscribing to TelemetryData\n")

	physicalInterfaceEvents := make(chan subscribers.PhysicalInterfaceEvent)
	loopbackInterfaceEvents := make(chan subscribers.LoopbackInterfaceEvent)
	subscribers.SubscribeToPhysicalInterfaceEvents(physicalInterfaceEvents)
	subscribers.SubscribeToLoopbackInterfaceEvents(loopbackInterfaceEvents)
	defer func() {
		subscribers.UnsubscribeFromPhysicalInterfaceEvents(physicalInterfaceEvents)
		subscribers.UnsubscribeFromLoopbackInterfaceEvents(loopbackInterfaceEvents)
	}()

	for {
		select {
		case event := <-physicalInterfaceEvents:
			if len(subscription.Ipv4Addresses) == 0 || helpers.IsInSlice(subscription.Ipv4Addresses, event.Ipv4Address) {
				response := convertPhysicalInterface(event, subscription.PropertyNames)
				err := responseStream.Send(&response)
				if err != nil {
					return err
				}
			}
		case event := <-loopbackInterfaceEvents:
			if len(subscription.Ipv4Addresses) == 0 || helpers.IsInSlice(subscription.Ipv4Addresses, event.Ipv4Address) {
				response := convertLoopbackInterface(event, subscription.PropertyNames)
				err := responseStream.Send(&response)
				if err != nil {
					return err
				}
			}
		}
	}
}

func (s *pushServiceServer) SubscribeToDataRate(subscription *DataRateSubscription, responseStream PushService_SubscribeToDataRateServer) error {
	log.Printf("SR-App subscribing to DataRate\n")

	physicalInterfaceEvents := make(chan subscribers.PhysicalInterfaceEvent)
	subscribers.SubscribeToPhysicalInterfaceEvents(physicalInterfaceEvents)
	defer func() {
		//TODO: defer is only called if client is exited (using eg ctrl + C) but not if context gets cancelled
		subscribers.UnsubscribeFromPhysicalInterfaceEvents(physicalInterfaceEvents)
	}()

	//TODO: check if stream was canceled from client side
	for {
		event := <-physicalInterfaceEvents
		if subscription.Ipv4Address == event.Ipv4Address {
			response := DataRate{DataRate: event.DataRate}
			err := responseStream.Send(&response)
			if err != nil {
				return err
			}
		}

	}
}
