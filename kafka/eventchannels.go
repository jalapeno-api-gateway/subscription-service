package kafka

var LsNodeEvents = make(chan KafkaEventMessage)
var LsLinkEvents = make(chan KafkaEventMessage)
var PhysicalInterfaceEvents = make(chan PhysicalInterfaceEventMessage)
var LoopbackInterfaceEvents = make(chan LoopbackInterfaceEventMessage)