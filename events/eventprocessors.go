package events

import (
	"context"
	"log"

	"gitlab.ost.ch/ins/jalapeno-api/push-service/arangodb"
	"gitlab.ost.ch/ins/jalapeno-api/push-service/kafka"
	"gitlab.ost.ch/ins/jalapeno-api/push-service/subscribers"
)

func StartEventProcessing() {
	for {
		select {
		case event := <-kafka.LsNodeEvents:
			handleLsNodeEvent(event)
		case event := <-kafka.LsLinkEvents:
			handleLsLinkEvent(event)
		case event := <-kafka.PhysicalInterfaceEvents:
			handlePhysicalInterfaceEvent(event)
		case event := <-kafka.LoopbackInterfaceEvents:
			handleLoopbackInterfaceEvent(event)
		}
	}
}

func handleLsNodeEvent(event kafka.KafkaEventMessage) {
	ctx := context.Background()
	updatedNode := arangodb.LsNodeDocument{}

	log.Printf("LsNode [%s]: %s\n", event.Action, event.Key)
	if event.Action != "del" {
		updatedNode = arangodb.FetchLsNode(ctx, event.Key)
	}

	nodeEvent := subscribers.LsNodeEvent{Action: event.Action, Key: event.Key, LsNodeDocument: updatedNode}
	subscribers.NotifyLsNodeSubscribers(nodeEvent)
}

func handleLsLinkEvent(event kafka.KafkaEventMessage) {
	ctx := context.Background()
	updatedLink := arangodb.LsLinkDocument{}

	log.Printf("LsLink [%s]: %s\n", event.Action, event.Key)
	if event.Action != "del" {
		updatedLink = arangodb.FetchLsLink(ctx, event.Key)
	}

	linkEvent := subscribers.LsLinkEvent{Action: event.Action, Key: event.Key, LsLinkDocument: updatedLink}
	subscribers.NotifyLsLinkSubscribers(linkEvent)
}

func handlePhysicalInterfaceEvent(event kafka.PhysicalInterfaceEventMessage) {
	phyiscalInterfaceEvent := subscribers.PhysicalInterfaceEvent{
		Ipv4Address:     event.IpAddress,
		DataRate:        event.DataRate,
		PacketsSent:     event.PacketsSent,
		PacketsReceived: event.PacketsReceived,
	}
	subscribers.NotifyPhysicalInterfaceSubscribers(phyiscalInterfaceEvent)
}

func handleLoopbackInterfaceEvent(event kafka.LoopbackInterfaceEventMessage) {
	loopbackInterfaceEvent := subscribers.LoopbackInterfaceEvent{
		Ipv4Address:     			event.IpAddress,
		State:        				event.State,
		LastStateTransitionTime: 	event.LastStateTransitionTime,
	}
	subscribers.NotifyLoopbackInterfaceSubscribers(loopbackInterfaceEvent)
}