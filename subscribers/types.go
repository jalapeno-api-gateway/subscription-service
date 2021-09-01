package subscribers

import (
	"gitlab.ost.ch/ins/jalapeno-api/push-service/arangodb"
	"gitlab.ost.ch/ins/jalapeno-api/push-service/kafka"
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

type DataRateEvent struct {
	Key      string
	DataRate kafka.DataRate
}

type TelemetryEvent struct {
	Key                  string
	DataRate             int64
	TotalPacketsSent     int64
	TotalPacketsReceived int64
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

type dataRateSubscriberUpdate struct {
	Action        Action
	UpdateChannel chan DataRateEvent
}

type telemetrySubscriberUpdate struct {
	Action        Action
	UpdateChannel chan TelemetryEvent
}
