package kafka

import (
	"context"

	"gitlab.ost.ch/ins/jalapeno-api/push-service/pubsub"
	"gitlab.ost.ch/ins/jalapeno-api/push-service/arangodb"
	"gitlab.ost.ch/ins/jalapeno-api/push-service/model"
)

func handleTopologyEvent(msg KafkaEventMessage, eventType model.EventType) {
	ctx := context.Background()
	var document interface{}
	if msg.Action != "del" {
		document = fetchDocument(ctx, msg.Key, eventType)
	}
	event := model.TopologyEvent{Action: msg.Action, Key: msg.Key, Document: document}
	publishTopologyEvent(event, eventType)
}

func fetchDocument(ctx context.Context, key string, eventType model.EventType) interface{} {
	switch eventType {
		case model.LsNodeEvent: return arangodb.FetchLsNode(ctx, key)
		case model.LsLinkEvent: return arangodb.FetchLsLink(ctx, key)
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
