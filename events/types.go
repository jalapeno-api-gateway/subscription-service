package events

import protocol "github.com/influxdata/line-protocol"

type TopologyEvent struct {
	Action   string
	Key      string
	Document interface{}
}

type TelemetryEvent struct {
	Measurement string
	Metric      protocol.Metric
}
