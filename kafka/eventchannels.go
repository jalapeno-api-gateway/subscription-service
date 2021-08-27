package kafka

var LsNodeEvents = make(chan KafkaEventMessage)
var LsLinkEvents = make(chan KafkaEventMessage)
var TelemetryDataRateEvents = make(chan KafkaTelemetryDataRateEventMessage)

//Add more different TelemtryXEvents channel for other attributes
