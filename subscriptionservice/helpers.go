package subscriptionservice

import (
	"context"
	"fmt"

	protocol "github.com/influxdata/line-protocol"
	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"github.com/jalapeno-api-gateway/subscription-service/helpers"
	"google.golang.org/grpc/peer"
)

func getClientIp(ctx context.Context) string {
	p, status := peer.FromContext(ctx)
	if status {
		return p.Addr.String()
	}
	return ""
}

func isSubscribed(metric protocol.Metric, stringFilters []*jagw.StringFilter) bool {
	for _, filter := range stringFilters {
		if !doesStringFilterMatch(metric, filter) {
			return false
		}
	}
	return true
}

func doesStringFilterMatch(metric protocol.Metric, filter *jagw.StringFilter) bool {
	value, err := helpers.GetValueByPropertyName(metric, *filter.Property)
	if err != nil {
		return false
	}

	valueAsString := fmt.Sprintf("%v", value)
	if valueAsString == *filter.Value {
		return true
	} else {
		return false
	}
}

