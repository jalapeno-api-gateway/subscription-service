package subscriptionservice

import (
	"github.com/jalapeno-api-gateway/arangodb-adapter/arango"
	"github.com/jalapeno-api-gateway/subscription-service/events"
	"google.golang.org/protobuf/proto"
)

func convertLSNodeEvent(event events.TopologyEvent) *LSNodeEvent {
	document := event.Document.(arango.LSNode)

	response := &LSNodeEvent{
		Action: proto.String(event.Action),
		Key:    proto.String(event.Key),
	}

	if event.Action == "del" {
		return response
	}

	response.LsNode = &LSNode{
		Key: proto.String(document.Key),
		ID: proto.String(document.ID),
		RouterHash: proto.String(document.RouterHash),
		DomainID: proto.Int64(document.DomainID),
		RouterIP: proto.String(document.RouterIP),
		PeerHash: proto.String(document.PeerHash),
		PeerIP: proto.String(document.PeerIP),
		PeerASN: proto.Int32(document.PeerASN),
		Timestamp: proto.String(document.Timestamp),
		IGPRouterID: proto.String(document.IGPRouterID),
		ASN: proto.Uint32(document.ASN),
		MTID: convertMTIDSlice(document.MTID),
		AreaID: proto.String(document.AreaID),
		Protocol: proto.String(document.Protocol),
		ProtocolID: proto.Uint32(uint32(document.ProtocolID)),
		Name: proto.String(document.Name),
		IsPrepolicy: proto.Bool(document.IsPrepolicy),
		IsAdjRIBIn: proto.Bool(document.IsAdjRIBIn),
	}

	return response
}

func convertLSLinkEvent(event events.TopologyEvent) *LSLinkEvent {
	document := event.Document.(arango.LSLink)

	response := &LSLinkEvent{
		Action: proto.String(event.Action),
		Key:    proto.String(event.Key),
	}

	if event.Action == "del" {
		return response
	}

	response.LsLink = &LSLink{
		Key: proto.String(document.Key),
		ID: proto.String(document.ID),
		RouterHash: proto.String(document.RouterHash),
		RouterIP: proto.String(document.RouterIP),
		DomainID: proto.Int64(document.DomainID),
		PeerHash: proto.String(document.PeerHash),
		PeerIP: proto.String(document.PeerIP),
		PeerASN: proto.Int32(document.PeerASN),
		Timestamp: proto.String(document.Timestamp),
		IGPRouterID: proto.String(document.IGPRouterID),
		Protocol: proto.String(document.Protocol),
		AreaID: proto.String(document.AreaID),
		Nexthop: proto.String(document.Nexthop),
		MTID: convertMTID(document.MTID),
		LocalLinkIP: proto.String(document.LocalLinkIP),
		RemoteLinkIP: proto.String(document.RemoteLinkIP),
		IGPMetric: proto.Uint32(document.IGPMetric),
		RemoteNodeHash: proto.String(document.RemoteNodeHash),
		LocalNodeHash: proto.String(document.LocalNodeHash),
		RemoteIGPRouterID: proto.String(document.RemoteIGPRouterID),
	}

	return response
}

func convertLSPrefixEvent(event events.TopologyEvent) *LSPrefixEvent {
	document := event.Document.(arango.LSPrefix)

	response := &LSPrefixEvent{
		Action: proto.String(event.Action),
		Key:    proto.String(event.Key),
	}

	if event.Action == "del" {
		return response
	}

	response.LsPrefix = &LSPrefix{
		Key: proto.String(document.Key),
		ID: proto.String(document.ID),
		RouterHash: proto.String(document.RouterHash),
		RouterIP: proto.String(document.RouterIP),
		DomainID: proto.Int64(document.DomainID),
		PeerHash: proto.String(document.PeerHash),
		PeerIP: proto.String(document.PeerIP),
		PeerASN: proto.Int32(document.PeerASN),
		Timestamp: proto.String(document.Timestamp),
		IGPRouterID: proto.String(document.IGPRouterID),
		Protocol: proto.String(document.Protocol),
		AreaID: proto.String(document.AreaID),
		Nexthop: proto.String(document.Nexthop),
		LocalNodeHash: proto.String(document.LocalNodeHash),
		MTID: convertMTID(document.MTID),
		Prefix: proto.String(document.Prefix),
		PrefixLen: proto.Int32(document.PrefixLen),
		PrefixMetric: proto.Uint32(document.PrefixMetric),
		IsPrepolicy: proto.Bool(document.IsPrepolicy),
		IsAdjRIBIn: proto.Bool(document.IsAdjRIBIn),
	}

	return response
}

func convertLSSRv6SIDEvent(event events.TopologyEvent) *LSSRv6SIDEvent {
	document := event.Document.(arango.LSSRv6SID)

	response := &LSSRv6SIDEvent{
		Action: proto.String(event.Action),
		Key:    proto.String(event.Key),
	}

	if event.Action == "del" {
		return response
	}

	response.LsSRv6SID = &LSSRv6SID{
		Key: proto.String(document.Key),
		ID: proto.String(document.ID),
		RouterHash: proto.String(document.RouterHash),
		RouterIP: proto.String(document.RouterIP),
		DomainID: proto.Int64(document.DomainID),
		PeerHash: proto.String(document.PeerHash),
		PeerIP: proto.String(document.PeerIP),
		PeerASN: proto.Int32(document.PeerASN),
		Timestamp: proto.String(document.Timestamp),
		IGPRouterID: proto.String(document.IGPRouterID),
		LocalNodeASN: proto.Uint32(document.LocalNodeASN),
		Protocol: proto.String(document.Protocol),
		Nexthop: proto.String(document.Nexthop),
		LocalNodeHash: proto.String(document.LocalNodeHash),
		MTID: convertMTID(document.MTID),
		IGPFlags: proto.Uint32(uint32(document.IGPFlags)),
		IsPrepolicy: proto.Bool(document.IsPrepolicy),
		IsAdjRIBIn: proto.Bool(document.IsAdjRIBIn),
		SRv6SID: proto.String(document.SRv6SID),
	}

	return response
}

func convertMTIDSlice(documents []*arango.MultiTopologyIdentifier) []*MultiTopologyIdentifier {
	mtids := []*MultiTopologyIdentifier{}
	for _, doc := range documents {
		mtids = append(mtids, convertMTID(doc))
	}
	return mtids
}

func convertMTID(doc *arango.MultiTopologyIdentifier) *MultiTopologyIdentifier {
	return &MultiTopologyIdentifier{
		OFlag: proto.Bool(doc.OFlag),
		AFlag: proto.Bool(doc.AFlag),
		MTID: proto.Uint32(uint32(doc.MTID)),
	}
}