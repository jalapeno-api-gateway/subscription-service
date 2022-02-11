package subscriptionservice

import (
	"context"

	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"google.golang.org/grpc/peer"
)

func getClientIp(ctx context.Context) string {
	p, status := peer.FromContext(ctx)
	if status {
		return p.Addr.String()
	}
	return ""
}

func isSubscribed(interfaceIds []*jagw.InterfaceIdentifier, hostname string, linkId int32) bool {
	for _, interfaceId := range interfaceIds {
		if *interfaceId.Hostname == hostname && *interfaceId.LinkId == linkId {
			return true
		}
	}
	return false
}