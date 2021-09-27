package kafka

import (
	"context"

	"github.com/Jalapeno-API-Gateway/subscription-service/pubsub"
	"github.com/jalapeno-api-gateway/arangodb-adapter/arango"
	"github.com/Jalapeno-API-Gateway/subscription-service/model"
)

func handleTopologyEvent(msg KafkaEventMessage, eventType model.EventType) {
	ctx := context.Background()
	document := fetchDocument(ctx, msg, eventType)
	event := model.TopologyEvent{Action: msg.Action, Key: msg.Key, Document: document}
	publishTopologyEvent(event, eventType)
}

func fetchDocument(ctx context.Context, msg KafkaEventMessage, eventType model.EventType) interface{} {
	if msg.Action == "del" {
		switch eventType {
			case model.LsNodeEvent: return arango.LsNodeDocument{}
			case model.LsLinkEvent: return arango.LsLinkDocument{}
			default: return nil
		}
	}

	switch eventType {
		case model.LsNodeEvent: return arango.FetchLsNode(ctx, msg.Key)
		case model.LsLinkEvent: return arango.FetchLsLink(ctx, msg.Key)
		default: return nil
	}
}

func publishTopologyEvent(event model.TopologyEvent, documentType model.EventType) {
	switch documentType {
		case model.LsNodeEvent: pubsub.LsNodeTopic.Publish(event)
		case model.LsLinkEvent: pubsub.LsLinkTopic.Publish(event)
	}
}

func handleTelemetryEvent(telemetryString string) {
	if !containsIpAddress(telemetryString) { // Only telemetry events containing an IP-Address are supported (they are used as identifiers)
		return
	}

	if isLoopbackEvent(telemetryString) {
		event := createLoopbackInterfaceEvent(telemetryString)
		pubsub.LoopbackInterfaceTopic.Publish(event)
	} else {
		event := createPhysicalInterfaceEvent(telemetryString)
		pubsub.PhysicalInterfaceTopic.Publish(event)
	}
}
