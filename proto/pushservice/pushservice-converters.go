package pushservice

import (
	"gitlab.ost.ch/ins/jalapeno-api/push-service/arangodb"
	"gitlab.ost.ch/ins/jalapeno-api/push-service/model"
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
	document := event.Document.(arangodb.LsNodeDocument)

	lsNode := &LsNode{
		Key:      document.Key,
		Name:     document.Name,
		Asn:      document.Asn,
		RouterIp: document.Router_ip,
	}

	return &LsNodeEvent{
		Action: event.Action,
		Key:    event.Key,
		LsNode: lsNode,
	}
}

func convertLsLinkEvent(event model.TopologyEvent) *LsLinkEvent {
	document := event.Document.(arangodb.LsLinkDocument)

	lsLink := &LsLink{
		Key:          document.Key,
		RouterIp:     document.Router_ip,
		PeerIp:       document.Peer_ip,
		LocalLinkIp:  document.LocalLink_ip,
		RemoteLinkIp: document.RemoteLink_ip,
		IgpMetric:    int32(document.Igp_metric),
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
			case DataRateProperty: telemetryEvent.DataRate = event.DataRate
			case PacketsSentProperty: telemetryEvent.PacketsSent = event.PacketsSent
			case PacketsReceivedProperty: telemetryEvent.PacketsReceived = event.PacketsReceived
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
			case StateProperty: telemetryEvent.State = event.State
			case LastStateTransitionTimeProperty: telemetryEvent.LastStateTransitionTime = event.LastStateTransitionTime
		}
	}
	
	return &telemetryEvent
}