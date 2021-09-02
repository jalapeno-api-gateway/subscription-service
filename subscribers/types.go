package subscribers

import (
	"gitlab.ost.ch/ins/jalapeno-api/push-service/arangodb"
)

//
// ---> ACTIONS <---
//

type Action string

const (
	Subscribe   Action = "subscribe"
	Unsubscribe Action = "unsubscribe"
)

//
// ---> EVENTS <---
//

type LsNodeEvent struct {
	Action         string
	Key            string
	LsNodeDocument arangodb.LsNodeDocument
}

type LsLinkEvent struct {
	Action         string
	Key            string
	LsLinkDocument arangodb.LsLinkDocument
}

type PhysicalInterfaceEvent struct {
	Ipv4Address		string
	DataRate        int64
	PacketsSent     int64
	PacketsReceived int64
}

type LoopbackInterfaceEvent struct {
	Ipv4Address					string
	State           			string
	LastStateTransitionTime     int64
}

//
// ---> SUBSCRIBER UPDATES <---
//

type lsNodeSubscriberUpdate struct {
	Action        Action
	UpdateChannel chan LsNodeEvent
}

type lsLinkSubscriberUpdate struct {
	Action        Action
	UpdateChannel chan LsLinkEvent
}

type physicalInterfaceSubscriberUpdate struct {
	Action        Action
	UpdateChannel chan PhysicalInterfaceEvent
}

type loopbackInterfaceSubscriberUpdate struct {
	Action        Action
	UpdateChannel chan LoopbackInterfaceEvent
}
