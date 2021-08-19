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

func (s *pushServiceServer) SubscribeToDataRates(subscription *DataRateSubscription, responseStream PushService_SubscribeToDataRatesServer) error {
	log.Printf("SR-App subscribing to DataRates\n")

	log.Printf("Subscription to DataRates ended")
	return nil
}

func (s *pushServiceServer) SubscribeToLsNodes(subscription *LsNodeSubscription, responseStream PushService_SubscribeToLsNodesServer) error {
	log.Printf("SR-App subscribing to LsNodes\n")

	events := make(chan subscribers.LsNodeEvent)
	subscribers.SubscribeToLsNodeEvents(events)
	defer func() {
		subscribers.UnSubscribeFromLsNodeEvents(events)
	}()

	for {
		event := <- events
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
		event := <- events
		if len(subscription.Keys) == 0 || helpers.IsInSlice(subscription.Keys, event.Key) {
			response := convertToGrpcLsLinkEvent(event)
			err := responseStream.Send(&response)
			if err != nil {
				return err
			}
		}
	}
}