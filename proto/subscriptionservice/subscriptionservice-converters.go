package subscriptionservice

import (
	"github.com/jalapeno-api-gateway/arangodb-adapter/arango"
	"github.com/Jalapeno-API-Gateway/subscription-service/model"
	"google.golang.org/protobuf/proto"
)

const (
	DataRateProperty = "DataRate"
	PacketsSentProperty = "PacketsSent"
	PacketsReceivedProperty = "PacketsReceived"
	StateProperty = "State"
	LastStateTransitionTimeProperty = "LastStateTransitionTime"
)

var allPhysicalInterfaceProperties = []string{
	DataRateProperty,
	PacketsSentProperty,
	PacketsReceivedProperty,
}

var allLoopbackInterfaceProperties = []string{
	StateProperty,
	LastStateTransitionTimeProperty,
}

func convertLsNodeEvent(event model.TopologyEvent) *LsNodeEvent {
	document := event.Document.(arango.LsNodeDocument)

	lsNode := &LsNode{
		Key:      document.Key,
		Name:     proto.String(document.Name),
		Asn:      proto.Int32(document.Asn),
		RouterIp: proto.String(document.Router_ip),
	}

	return &LsNodeEvent{
		Action: event.Action,
		Key:    event.Key,
		LsNode: lsNode,
	}
}

func convertLsLinkEvent(event model.TopologyEvent) *LsLinkEvent {
	document := event.Document.(arango.LsLinkDocument)

	lsLink := &LsLink{
		Key:          document.Key,
		RouterIp:     proto.String(document.Router_ip),
		PeerIp:       proto.String(document.Peer_ip),
		LocalLinkIp:  proto.String(document.LocalLink_ip),
		RemoteLinkIp: proto.String(document.RemoteLink_ip),
		IgpMetric:    proto.Int32(int32(document.Igp_metric)),
	}

	return &LsLinkEvent{
		Action: event.Action,
		Key:    event.Key,
		LsLink: lsLink,
	}
}

func convertPhysicalInterfaceEvent(event model.PhysicalInterfaceEvent, propertyNames []string) *TelemetryEvent {
	if len(propertyNames) == 0 { // If no propertyNames were provided, all Properties are returned to the SR-App
		propertyNames = allPhysicalInterfaceProperties
	}

	telemetryEvent := TelemetryEvent{
		Ipv4Address: event.Ipv4Address,
	}

	for _, propertyName := range propertyNames {
		switch propertyName {
			case DataRateProperty: telemetryEvent.DataRate = proto.Int64(event.DataRate)
			case PacketsSentProperty: telemetryEvent.PacketsSent = proto.Int64(event.PacketsSent)
			case PacketsReceivedProperty: telemetryEvent.PacketsReceived = proto.Int64(event.PacketsReceived)
		}
	}
	
	return &telemetryEvent
}

func convertLoopbackInterfaceEvent(event model.LoopbackInterfaceEvent, propertyNames []string) *TelemetryEvent {
	if len(propertyNames) == 0 { // If no propertyNames were provided, all Properties are returned to the SR-App
		propertyNames = allPhysicalInterfaceProperties
	}

	telemetryEvent := TelemetryEvent{
		Ipv4Address: event.Ipv4Address,
	}

	for _, propertyName := range propertyNames {
		switch propertyName {
			case StateProperty: telemetryEvent.State = proto.String(event.State)
			case LastStateTransitionTimeProperty: telemetryEvent.LastStateTransitionTime = proto.Int64(event.LastStateTransitionTime)
		}
	}
	
	return &telemetryEvent
}