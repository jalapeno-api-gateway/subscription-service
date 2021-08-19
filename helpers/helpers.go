package helpers

import (
	"os"
	"os/signal"
)

func IsInSlice(slice []string, value string) bool {
	for _, elem := range slice {
        if elem == value {
            return true
        }
    }
    return false
}

func WatchInterruptSignals() chan os.Signal {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
    return signals
}