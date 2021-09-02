package pushservice

import (
	"encoding/json"
	"errors"

	"gitlab.ost.ch/ins/jalapeno-api/push-service/arangodb"
	"gitlab.ost.ch/ins/jalapeno-api/push-service/subscribers"
)

const (
	dataRate = "DataRate"
	packetsSent = "PacketsSent"
	packetsReceived = "PacketsReceived"
	state = "State"
	lastStateTransitionTime = "LastStateTransitionTime"
)

func convertToGrpcLsNodeEvent(event subscribers.LsNodeEvent) LsNodeEvent {
	lsNode := convertToGrpcLsNode(event.LsNodeDocument)
	return LsNodeEvent{
		Action: event.Action,
		Key:    event.Key,
		LsNode: &lsNode,
	}
}

func convertToGrpcLsLinkEvent(event subscribers.LsLinkEvent) LsLinkEvent {
	lsLink := convertToGrpcLsLink(event.LsLinkDocument)
	return LsLinkEvent{
		Action: event.Action,
		Key:    event.Key,
		LsLink: &lsLink,
	}
}

func convertToGrpcLsNode(nodeDocument arangodb.LsNodeDocument) LsNode {
	return LsNode{
		Key:      nodeDocument.Key,
		Name:     nodeDocument.Name,
		Asn:      nodeDocument.Asn,
		RouterIp: nodeDocument.Router_ip,
	}
}

func convertToGrpcLsLink(linkDocument arangodb.LsLinkDocument) LsLink {
	return LsLink{
		Key:          linkDocument.Key,
		RouterIp:     linkDocument.Router_ip,
		PeerIp:       linkDocument.Peer_ip,
		LocalLinkIp:  linkDocument.LocalLink_ip,
		RemoteLinkIp: linkDocument.RemoteLink_ip,
		IgpMetric:    int32(linkDocument.Igp_metric),
	}
}

func convertLoopbackInterface(event subscribers.LoopbackInterfaceEvent, propertyNames []string) TelemetryEvent {
	data := []*TelemetryData{}

	if len(propertyNames) == 0 {
		propertyNames = []string{
			state,
			lastStateTransitionTime,
		}
	}

	for _, propertyName := range propertyNames {
		telemetryData, err := getLoopbackInterfaceData(event, propertyName)
		if err == nil {
			data = append(data, &telemetryData)
		}
	}

	return TelemetryEvent{
		Ipv4Address: event.Ipv4Address,
		Data: data,
	}
}

func convertPhysicalInterface(event subscribers.PhysicalInterfaceEvent, propertyNames []string) TelemetryEvent {
	data := []*TelemetryData{}

	if len(propertyNames) == 0 {
		propertyNames = []string{
			dataRate,
			packetsSent,
			packetsReceived,
		}
	}

	for _, propertyName := range propertyNames {
		telemetryData, err := getPhysicalInterfaceData(event, propertyName)
		if err == nil {
			data = append(data, &telemetryData)
		}
	}

	return TelemetryEvent{
		Ipv4Address: event.Ipv4Address,
		Data: data,
	}
}

func getLoopbackInterfaceData(event subscribers.LoopbackInterfaceEvent, propertyName string) (TelemetryData, error) {
	value := []byte{}
	var err error

	switch propertyName {
		case state:
			value, err = json.Marshal(event.State)
		case lastStateTransitionTime:
			value, err = json.Marshal(event.LastStateTransitionTime)
		default:
			err = errors.New("Invalid property name!")
	}

	if err != nil {
		return TelemetryData{}, err
	}
	return TelemetryData{PropertyName: propertyName, Value: value}, nil
}


func getPhysicalInterfaceData(event subscribers.PhysicalInterfaceEvent, propertyName string) (TelemetryData, error) {
	value := []byte{}
	var err error

	switch propertyName {
		case dataRate:
			value, err = json.Marshal(event.DataRate)
		case packetsSent:
			value, err = json.Marshal(event.PacketsSent)
		case packetsReceived:
			value, err = json.Marshal(event.PacketsReceived)
		default:
			err = errors.New("Invalid property name!")
	}

	if err != nil {
		return TelemetryData{}, err
	}
	return TelemetryData{PropertyName: propertyName, Value: value}, nil
}
