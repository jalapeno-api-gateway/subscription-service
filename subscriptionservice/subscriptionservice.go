package subscriptionservice

import (
	"context"

	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"github.com/jalapeno-api-gateway/subscription-service/events"
	"github.com/jalapeno-api-gateway/subscription-service/helpers"
	"github.com/jalapeno-api-gateway/subscription-service/pubsub"
	"github.com/sirupsen/logrus"
)

type subscriptionServiceServer struct {
	jagw.UnimplementedSubscriptionServiceServer
}

func NewServer() *subscriptionServiceServer {
	s := &subscriptionServiceServer{}
	return s
}

func (s *subscriptionServiceServer) SubscribeToLsNodes(subscription *jagw.TopologySubscription, responseStream jagw.SubscriptionService_SubscribeToLsNodesServer) error {
	logger := logrus.WithFields(logrus.Fields{"clientIp": getClientIp(responseStream.Context()), "grpcFunction": "SubscribeToLsNodes"})
	logger.Debug("Incoming request.")

	cctx, cancel := context.WithCancel(context.Background())
	sub := pubsub.LsNodeTopic.Subscribe(logger)
	defer func() {
		sub.Unsubscribe(logger)
	}()
	
	sub.Receive(cctx, func(msg *interface{}) {
		event := (*msg).(events.TopologyEvent)
		logger = logger.WithFields(logrus.Fields{"key": event.Key, "Action": event.Action})
		logger.Debug("Subscription received new message.")

		if len(subscription.Keys) == 0 || helpers.IsInSlice(subscription.Keys, event.Key) {
			response := convertLsNodeEvent(event)
			logger.Debug("Sending response through gRPC stream.")
			err := responseStream.Send(response)
			if err != nil {
				logger.WithError(err).Error("Stream is aborting due to an error.")
				cancel()
			}
		}
	})

	return nil
}

func (s *subscriptionServiceServer) SubscribeToLsLinks(subscription *jagw.TopologySubscription, responseStream jagw.SubscriptionService_SubscribeToLsLinksServer) error {
	logger := logrus.WithFields(logrus.Fields{"clientIp": getClientIp(responseStream.Context()), "grpcFunction": "SubscribeToLsLinks"})
	logger.Debug("Incoming request.")

	cctx, cancel := context.WithCancel(context.Background())
	sub := pubsub.LsLinkTopic.Subscribe(logger)
	defer func() {
		sub.Unsubscribe(logger)
	}()

	sub.Receive(cctx, func(msg *interface{}) {
		event := (*msg).(events.TopologyEvent)
		logger = logger.WithFields(logrus.Fields{"key": event.Key, "Action": event.Action})
		logger.Debug("Subscription received new message.")

		if len(subscription.Keys) == 0 || helpers.IsInSlice(subscription.Keys, event.Key) {
			response := convertLsLinkEvent(event)
			logger.Debug("Sending response through gRPC stream.")
			err := responseStream.Send(response)
			if err != nil {
				logger.WithError(err).Error("Stream is aborting due to an error.")
				cancel()
			}
		}
	})

	return nil
}

func (s *subscriptionServiceServer) SubscribeToLsPrefixes(subscription *jagw.TopologySubscription, responseStream jagw.SubscriptionService_SubscribeToLsPrefixesServer) error {
	logger := logrus.WithFields(logrus.Fields{"clientIp": getClientIp(responseStream.Context()), "grpcFunction": "SubscribeToLsPrefixes"})
	logger.Debug("Incoming request.")

	cctx, cancel := context.WithCancel(context.Background())
	sub := pubsub.LsPrefixTopic.Subscribe(logger)
	defer func() {
		sub.Unsubscribe(logger)
	}()

	sub.Receive(cctx, func(msg *interface{}) {
		event := (*msg).(events.TopologyEvent)
		logger = logger.WithFields(logrus.Fields{"key": event.Key, "Action": event.Action})
		logger.Debug("Subscription received new message.")

		if len(subscription.Keys) == 0 || helpers.IsInSlice(subscription.Keys, event.Key) {
			response := convertLsPrefixEvent(event)
			logger.Debug("Sending response through gRPC stream.")
			err := responseStream.Send(response)
			if err != nil {
				logger.WithError(err).Error("Stream is aborting due to an error.")
				cancel()
			}
		}
	})

	return nil
}

func (s *subscriptionServiceServer) SubscribeToLsSrv6Sids(subscription *jagw.TopologySubscription, responseStream jagw.SubscriptionService_SubscribeToLsSrv6SidsServer) error {
	logger := logrus.WithFields(logrus.Fields{"clientIp": getClientIp(responseStream.Context()), "grpcFunction": "SubscribeToLsSrv6Sids"})
	logger.Debug("Incoming request.")

	cctx, cancel := context.WithCancel(context.Background())
	sub := pubsub.LsSrv6SidTopic.Subscribe(logger)
	defer func() {
		sub.Unsubscribe(logger)
	}()

	sub.Receive(cctx, func(msg *interface{}) {
		event := (*msg).(events.TopologyEvent)
		logger = logger.WithFields(logrus.Fields{"key": event.Key, "Action": event.Action})
		logger.Debug("Subscription received new message.")

		if len(subscription.Keys) == 0 || helpers.IsInSlice(subscription.Keys, event.Key) {
			response := convertLsSrv6SidEvent(event)
			logger.Debug("Sending response through gRPC stream.")
			err := responseStream.Send(response)
			if err != nil {
				logger.WithError(err).Error("Stream is aborting due to an error.")
				cancel()
			}
		}
	})

	return nil
}

func (s *subscriptionServiceServer) SubscribeToLsNodeEdges(subscription *jagw.TopologySubscription, responseStream jagw.SubscriptionService_SubscribeToLsNodeEdgesServer) error {
	logger := logrus.WithFields(logrus.Fields{"clientIp": getClientIp(responseStream.Context()), "grpcFunction": "SubscribeToLsNodeEdges"})
	logger.Debug("Incoming request.")

	cctx, cancel := context.WithCancel(context.Background())
	sub := pubsub.LsNodeEdgeTopic.Subscribe(logger)
	defer func() {
		sub.Unsubscribe(logger)
	}()

	sub.Receive(cctx, func(msg *interface{}) {
		event := (*msg).(events.TopologyEvent)
		logger = logger.WithFields(logrus.Fields{"key": event.Key, "Action": event.Action})
		logger.Debug("Subscription received new message.")

		if len(subscription.Keys) == 0 || helpers.IsInSlice(subscription.Keys, event.Key) {
			response := convertLsNodeEdgeEvent(event)
			logger.Debug("Sending response through gRPC stream.")
			err := responseStream.Send(response)
			if err != nil {
				logger.WithError(err).Error("Stream is aborting due to an error.")
				cancel()
			}
		}
	})

	return nil
}

func (s *subscriptionServiceServer) SubscribeToTelemetryData(subscription *jagw.TelemetrySubscription, responseStream jagw.SubscriptionService_SubscribeToTelemetryDataServer) error {
	logger := logrus.WithFields(logrus.Fields{"clientIp": getClientIp(responseStream.Context()), "grpcFunction": "SubscribeToTelemetryData"})
	logger.Debug("Incoming request.")

	cctxA, cancelA := context.WithCancel(context.Background())
	cctxB, cancelB := context.WithCancel(context.Background())
	physicalInterfaceSubscription := pubsub.PhysicalInterfaceTopic.Subscribe(logger)
	loopbackInterfaceSubscription := pubsub.LoopbackInterfaceTopic.Subscribe(logger)
	
	defer func() {
		physicalInterfaceSubscription.Unsubscribe(logger)
		loopbackInterfaceSubscription.Unsubscribe(logger)
	}()

	go func() {
		physicalInterfaceSubscription.Receive(cctxA, func(msg *interface{}) {
			event := (*msg).(events.PhysicalInterfaceEvent)
			logger.Debug("Subscription received new message.")

			if len(subscription.InterfaceIds) == 0 || isSubscribed(subscription.InterfaceIds, event.Hostname, event.LinkID) {
				response := convertPhysicalInterfaceEvent(event, subscription.PropertyNames)
				logger.Debug("Sending response through gRPC stream.")
				err := responseStream.Send(response)
				if err != nil {
					logger.WithError(err).Error("Stream is aborting due to an error.")
					cancelA()
				}
			}
		})
	}()

	loopbackInterfaceSubscription.Receive(cctxB, func(msg *interface{}) {
		event := (*msg).(events.LoopbackInterfaceEvent)
		logger.Debug("Subscription received new message.")

		if len(subscription.InterfaceIds) == 0 || isSubscribed(subscription.InterfaceIds, event.Hostname, event.LinkID) {
			response := convertLoopbackInterfaceEvent(event, subscription.PropertyNames)
			logger.Debug("Sending response through gRPC stream.")
			err := responseStream.Send(response)
			if err != nil {
				logger.WithError(err).Error("Stream is aborting due to an error.")
				cancelB()
			}
		}
	})
	// TODO return status.Errorf(codes.)
	return nil
}
