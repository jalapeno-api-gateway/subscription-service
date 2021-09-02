package kafka

type KafkaEventMessage struct {
	TopicType int    `json:"TopicType,omitempty"`
	Key       string `json:"_key,omitempty"`
	Id        string `json:"_id,omitempty"`
	Action    string `json:"action,omitempty"`
}

type PhysicalInterfaceEventMessage struct {
	IpAddress		string
	DataRate		int64
	PacketsSent		int64
	PacketsReceived	int64
}

type LoopbackInterfaceEventMessage struct {
	IpAddress					string
	State						string
	LastStateTransitionTime		int64
}
