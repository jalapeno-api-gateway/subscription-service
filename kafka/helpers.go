package kafka

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/Shopify/sarama"
)

const (
	interfaceNameIdentifier = "interface_name"
	ipAddressIdentifier = "ip_information/ip_address"
	dataRateIdentifier = "data_rates/output_data_rate"
	packetsSentIdentifier = "interface_statistics/full_interface_stats/packets_sent"
	packetsReceivedIdentifier = "interface_statistics/full_interface_stats/packets_received"
	stateIdentifier = "state"
	lastStateTransitionTimeIdentifier = "last_state_transition_time"
)

func unmarshalKafkaMessage(msg *sarama.ConsumerMessage) KafkaEventMessage {
	var event KafkaEventMessage
	err := json.Unmarshal(msg.Value, &event)
	if err != nil {
		log.Fatalf("Could not unmarshal kafka message, %v", err)
	}
	return event
}

func createLoopbackInterfaceEvent(telemetryString string) LoopbackInterfaceEventMessage {
	ipAddress := getIpAddress(telemetryString)
	state := getState(telemetryString)
	lastStateTransitionTime := getLastStateTransitionTime(telemetryString)

	return LoopbackInterfaceEventMessage{
		IpAddress:       			ipAddress,
		State:        				state,
		LastStateTransitionTime:    int64(lastStateTransitionTime),
	}
}

func createPhysicalInterfaceEvent(telemetryString string) PhysicalInterfaceEventMessage {
	dataRate := getDataRate(telemetryString)
	ipAddress := getIpAddress(telemetryString)
	totalPacketsSent := getPacketsSent(telemetryString)
	totalPacketsReceived := getPacketsReceived(telemetryString)

	return PhysicalInterfaceEventMessage{
		IpAddress:       ipAddress,
		DataRate:        int64(dataRate),
		PacketsSent:     int64(totalPacketsSent),
		PacketsReceived: int64(totalPacketsReceived),
	}
}

func isLoopbackEvent(telemetryString string) bool {
	interfaceName := getInterfaceName(telemetryString)
	return strings.HasPrefix(interfaceName, "Loopback")
}

func containsIpAddress(telemetryString string) bool {
	return strings.Index(telemetryString, ipAddressIdentifier) != -1
}

func getInterfaceName(telemetryString string) string {
	return extractStringValue(telemetryString, interfaceNameIdentifier)
}

func getState(telemetryString string) string {
	untrimmedValue := extractStringValue(telemetryString, stateIdentifier)
	trimmedValue := strings.Trim(untrimmedValue, "\"") //every ip address contains a leading and tailing apostrophe
	return trimmedValue
}

func getIpAddress(telemetryString string) string {
	untrimmedValue := extractStringValue(telemetryString, ipAddressIdentifier)
	trimmedValue := strings.Trim(untrimmedValue, "\"") //every ip address contains a leading and tailing apostrophe
	return trimmedValue
}

func getLastStateTransitionTime(telemetryString string) int {
	return extractIntValue(telemetryString, lastStateTransitionTimeIdentifier)
}

func getDataRate(telemetryString string) int {
	return extractIntValue(telemetryString, dataRateIdentifier)
}

func getPacketsSent(telemetryString string) int {
	return extractIntValue(telemetryString, packetsSentIdentifier)
}

func getPacketsReceived(telemetryString string) int {
	return extractIntValue(telemetryString, packetsReceivedIdentifier)
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