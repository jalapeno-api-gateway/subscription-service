package kafka

import (
	"context"

	"github.com/jalapeno-api-gateway/subscription-service/pubsub"
	"github.com/jalapeno-api-gateway/arangodb-adapter/arango"
	"github.com/jalapeno-api-gateway/subscription-service/model"
	"github.com/jalapeno-api-gateway/model/class"
)

func handleTopologyEvent(msg KafkaEventMessage, className class.Class) {
	ctx := context.Background()
	document := fetchDocument(ctx, msg, className)
	event := model.TopologyEvent{Action: msg.Action, Key: msg.Key, Document: document}
	publishTopologyEvent(event, className)
}

func fetchDocument(ctx context.Context, msg KafkaEventMessage, className class.Class) interface{} {
	if msg.Action == "del" {
		switch className {
			case class.LSNode: return arango.LSNode{}
			case class.LSLink: return arango.LSLink{}
			default: return nil
		}
	}

	switch className {
		case class.LSNode: return arango.FetchLSNode(ctx, msg.Key)
		case class.LSLink: return arango.FetchLSLink(ctx, msg.Key)
		default: return nil
	}
}

func publishTopologyEvent(event model.TopologyEvent, className class.Class) {
	switch className {
		case class.LSNode: pubsub.LsNodeTopic.Publish(event)
		case class.LSLink: pubsub.LsLinkTopic.Publish(event)
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
