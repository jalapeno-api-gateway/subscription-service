package events

type TopologyEvent struct {
	Action   string
	Key      string
	Document interface{}
}

type PhysicalInterfaceEvent struct {
	Hostname string
	LinkID int32
	Ipv4Address		string
	DataRate        int64
	PacketsSent     int64
	PacketsReceived int64
}

type LoopbackInterfaceEvent struct {
	Hostname string
	LinkID int32
	Ipv4Address					string
	State           			string
	LastStateTransitionTime     int64
}