package model

type TopologyEvent struct {
	Action   string
	Key      string
	Document interface{}
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