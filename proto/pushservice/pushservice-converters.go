package pushservice

import (
	"encoding/json"
	"reflect"

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

func convertLsNodeEvent(event model.TopologyEvent) LsNodeEvent {
	document := event.Document.(arangodb.LsNodeDocument)

	lsNode := &LsNode{
		Key:      document.Key,
		Name:     document.Name,
		Asn:      document.Asn,
		RouterIp: document.Router_ip,
	}

	return LsNodeEvent{
		Action: event.Action,
		Key:    event.Key,
		LsNode: lsNode,
	}
}

func convertLsLinkEvent(event model.TopologyEvent) LsLinkEvent {
	document := event.Document.(arangodb.LsLinkDocument)

	lsLink := &LsLink{
		Key:          document.Key,
		RouterIp:     document.Router_ip,
		PeerIp:       document.Peer_ip,
		LocalLinkIp:  document.LocalLink_ip,
		RemoteLinkIp: document.RemoteLink_ip,
		IgpMetric:    int32(document.Igp_metric),
	}

	return LsLinkEvent{
		Action: event.Action,
		Key:    event.Key,
		LsLink: lsLink,
	}
}

func convertTelemetryEvent(event interface{}, propertyNames []string, eventType model.EventType) TelemetryEvent {
	data := []*TelemetryData{}

	if len(propertyNames) == 0 { // If no propertyNames were provided, all Properties are returned to the SR-App
		switch eventType {
			case model.PhysicalInterfaceTelemetryEvent: propertyNames = allPhysicalInterfaceProperties
			case model.LoopbackInterfaceTelemetryEvent: propertyNames = allLoopbackInterfaceProperties
		}
	}

	for _, propertyName := range propertyNames {
		telemetryData, err := getTelemetryData(event, propertyName, eventType)
		if err == nil {
			data = append(data, telemetryData)
		} // Ignore/Skip in case of error (SR-App provided invalid propertyName)
	}

	var ipv4Address string
	switch eventType {
		case model.PhysicalInterfaceTelemetryEvent: ipv4Address = event.(model.PhysicalInterfaceEvent).Ipv4Address
		case model.LoopbackInterfaceTelemetryEvent: ipv4Address = event.(model.LoopbackInterfaceEvent).Ipv4Address
	}

	return TelemetryEvent{
		Ipv4Address: ipv4Address,
		Data: data,
	}
}

func getTelemetryData(event interface{}, field string, eventType model.EventType) (*TelemetryData, error) {
	r := reflect.ValueOf(event)
	value := reflect.Indirect(r).FieldByName(field)
	bytes, err := json.Marshal(value.Interface())
	if err != nil {
		return &TelemetryData{}, err
	}
	return &TelemetryData{PropertyName: field, Value: bytes}, nil
}