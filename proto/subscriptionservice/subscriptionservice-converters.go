package subscriptionservice

import (
	"github.com/jalapeno-api-gateway/arangodb-adapter/arango"
	"github.com/jalapeno-api-gateway/subscription-service/model"
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
	document := event.Document.(arango.LSNode)

	lsNode := &LsNode{
		Key:      document.Key,
		Name:     proto.String(document.Name),
		Asn:      proto.Int32(int32(document.ASN)),
		RouterIp: proto.String(document.RouterIP),
	}

	return &LsNodeEvent{
		Action: event.Action,
		Key:    event.Key,
		LsNode: lsNode,
	}
}

func convertLsLinkEvent(event model.TopologyEvent) *LsLinkEvent {
	document := event.Document.(arango.LSLink)

	lsLink := &LsLink{
		Key:          document.Key,
		RouterIp:     proto.String(document.RouterIP),
		PeerIp:       proto.String(document.PeerIP),
		LocalLinkIp:  proto.String(document.LocalLinkIP),
		RemoteLinkIp: proto.String(document.RemoteLinkIP),
		IgpMetric:    proto.Int32(int32(document.IGPMetric)),
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