package pushservice

import (
	"gitlab.ost.ch/ins/jalapeno-api/push-service/arangodb"
	"gitlab.ost.ch/ins/jalapeno-api/push-service/influxdb"
	"gitlab.ost.ch/ins/jalapeno-api/push-service/subscribers"
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

func convertToGrpcDataRateEvent(event subscribers.DataRateEvent) DataRateEvent {
	dataRate := convertToGrpcDataRate(event.DataRate)
	return DataRateEvent{
		Key:      event.Key,
		DataRate: &dataRate,
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

func convertToGrpcDataRate(dataRate influxdb.DataRate) DataRate {
	return DataRate{
		Ipv4Address: dataRate.Ipv4Address,
		DataRate:    dataRate.DataRate,
	}
}
