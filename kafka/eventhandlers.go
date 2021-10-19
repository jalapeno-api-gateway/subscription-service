package kafka

import (
	"context"
	
	"github.com/jalapeno-api-gateway/jagw-core/arango"
	"github.com/jalapeno-api-gateway/jagw-core/model/class"
	"github.com/jalapeno-api-gateway/subscription-service/events"
	"github.com/jalapeno-api-gateway/subscription-service/pubsub"
)

func handleTopologyEvent(msg KafkaEventMessage, className class.Class) {
	ctx := context.Background()
	document := fetchDocument(ctx, msg, className)
	event := events.TopologyEvent{Action: msg.Action, Key: msg.Key, Document: document}
	publishTopologyEvent(event, className)
}

func fetchDocument(ctx context.Context, msg KafkaEventMessage, className class.Class) interface{} {
	if msg.Action == "del" {
		switch className {
			case class.LsNode: return arango.LSNode{}
			case class.LsLink: return arango.LSLink{}
			case class.LsPrefix: return arango.LSPrefix{}
			case class.LsSrv6Sid: return arango.LSSRv6SID{}
			case class.LsNodeEdge: return arango.LSNode_Edge{}
			default: return nil
		}
	}

	switch className {
		case class.LsNode: return arango.FetchLsNode(ctx, msg.Key)
		case class.LsLink: return arango.FetchLsLink(ctx, msg.Key)
		case class.LsPrefix: return arango.FetchLsPrefix(ctx, msg.Key)
		case class.LsSrv6Sid: return arango.FetchLsSrv6Sid(ctx, msg.Key)
		case class.LsNodeEdge: return arango.FetchLsNodeEdge(ctx, msg.Key)
		default: return nil
	}
}

func publishTopologyEvent(event events.TopologyEvent, className class.Class) {
	switch className {
		case class.LsNode: pubsub.LsNodeTopic.Publish(event)
		case class.LsLink: pubsub.LsLinkTopic.Publish(event)
		case class.LsPrefix: pubsub.LsPrefixTopic.Publish(event)
		case class.LsSrv6Sid: pubsub.LsSrv6SidTopic.Publish(event)
		case class.LsNodeEdge: pubsub.LsNodeEdgeTopic.Publish(event)
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
