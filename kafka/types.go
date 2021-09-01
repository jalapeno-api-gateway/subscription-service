package kafka

type KafkaEventMessage struct {
	TopicType int    `json:"TopicType,omitempty"`
	Key       string `json:"_key,omitempty"`
	Id        string `json:"_id,omitempty"`
	Action    string `json:"action,omitempty"`
}

type KafkaTelemetryDataRateEventMessage struct {
	IpAddress string
	DataRate  int64
}

type KafkaTelemetryEventMessage struct {
	IpAddress            string
	DataRate             int64
	TotalPacketsSent     int64
	TotalPacketsReceived int64
}

type DataRate struct {
	Ipv4Address string
	DataRate    int64
}
