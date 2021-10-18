package kafka

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/jalapeno-api-gateway/subscription-service/events"
)

const (
	HostNameIdentifier = "source"
	LinkIdIdentifier = "if_index"
	InterfaceNameIdentifier = "interface_name"
	IpAddressIdentifier = "ip_information/ip_address"
	DataRateIdentifier = "data_rates/output_data_rate"
	PacketsSentIdentifier = "interface_statistics/full_interface_stats/packets_sent"
	PacketsReceivedIdentifier = "interface_statistics/full_interface_stats/packets_received"
	StateIdentifier = "state"
	LastStateTransitionTimeIdentifier = "last_state_transition_time"
)

func unmarshalKafkaMessage(msg *sarama.ConsumerMessage) KafkaEventMessage {
	var event KafkaEventMessage
	err := json.Unmarshal(msg.Value, &event)
	if err != nil {
		log.Fatalf("Could not unmarshal kafka message, %v", err)
	}
	return event
}

func createLoopbackInterfaceEvent(telemetryString string) events.LoopbackInterfaceEvent {
	hostName := getHostName(telemetryString)
	linkId := getLinkId(telemetryString)
	ipv4Address := getIpAddress(telemetryString)
	state := getState(telemetryString)
	lastStateTransitionTime := getLastStateTransitionTime(telemetryString)

	return events.LoopbackInterfaceEvent {
		Hostname: 					hostName,
		LinkID: 					int32(linkId),
		Ipv4Address:       			ipv4Address,
		State:        				state,
		LastStateTransitionTime:    int64(lastStateTransitionTime),
	}
}

func createPhysicalInterfaceEvent(telemetryString string) events.PhysicalInterfaceEvent {
	hostName := getHostName(telemetryString)
	linkId := getLinkId(telemetryString)
	dataRate := getDataRate(telemetryString)
	ipv4Address := getIpAddress(telemetryString)
	totalPacketsSent := getPacketsSent(telemetryString)
	totalPacketsReceived := getPacketsReceived(telemetryString)

	return events.PhysicalInterfaceEvent {
		Hostname: 			hostName,
		LinkID: 			int32(linkId),
		Ipv4Address:     	ipv4Address,
		DataRate:        	int64(dataRate),
		PacketsSent:     	int64(totalPacketsSent),
		PacketsReceived: 	int64(totalPacketsReceived),
	}
}

func isLoopbackEvent(telemetryString string) bool {
	interfaceName := getInterfaceName(telemetryString)
	return strings.HasPrefix(interfaceName, "Loopback")
}

func containsIpAddress(telemetryString string) bool {
	return strings.Contains(telemetryString, IpAddressIdentifier)
}

func getHostName(telemetryString string) string {
	return extractStringValue(telemetryString, HostNameIdentifier)
}

func getLinkId(telemetryString string) int {
	return extractIntValue(telemetryString, LinkIdIdentifier)
}

func getInterfaceName(telemetryString string) string {
	return extractStringValue(telemetryString, InterfaceNameIdentifier)
}

func getState(telemetryString string) string {
	untrimmedValue := extractStringValue(telemetryString, StateIdentifier)
	trimmedValue := strings.Trim(untrimmedValue, "\"") //every ip address contains a leading and tailing apostrophe
	return trimmedValue
}

func getIpAddress(telemetryString string) string {
	untrimmedValue := extractStringValue(telemetryString, IpAddressIdentifier)
	trimmedValue := strings.Trim(untrimmedValue, "\"") //every ip address contains a leading and tailing apostrophe
	return trimmedValue
}

func getLastStateTransitionTime(telemetryString string) int {
	return extractIntValue(telemetryString, LastStateTransitionTimeIdentifier)
}

func getDataRate(telemetryString string) int {
	return extractIntValue(telemetryString, DataRateIdentifier)
}

func getPacketsSent(telemetryString string) int {
	return extractIntValue(telemetryString, PacketsSentIdentifier)
}

func getPacketsReceived(telemetryString string) int {
	return extractIntValue(telemetryString, PacketsReceivedIdentifier)
}

func extractIntValue(telemetryString string, propertyName string) int {
	untrimmedValue := extractStringValue(telemetryString, propertyName)
	trimmedValue := untrimmedValue[:len(untrimmedValue)-1] //every int value has the letter 'i' at the end
	value, err := strconv.Atoi(trimmedValue)
	if err != nil {
		log.Fatalf("Failed to convert string to int: %v", err)
	}
	return value
}

func extractStringValue(telemetryString string, propertyName string) string {
	index := strings.Index(telemetryString, propertyName)
	substring1 := telemetryString[index:]
	indexOfComma := strings.Index(substring1, ",")
	substring2 := substring1[:indexOfComma]
	split := strings.Split(substring2, "=")
	return split[1]
}