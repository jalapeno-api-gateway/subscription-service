package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/jalapeno-api-gateway/jagw-core/arango"
	"github.com/jalapeno-api-gateway/jagw-core/logger"
	"github.com/jalapeno-api-gateway/protorepo-jagw-go/jagw"
	"github.com/jalapeno-api-gateway/subscription-service/helpers"
	"github.com/jalapeno-api-gateway/subscription-service/kafka"
	"github.com/jalapeno-api-gateway/subscription-service/pubsub"
	"github.com/jalapeno-api-gateway/subscription-service/subscriptionservice"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	logger.Init(logrus.StandardLogger(), os.Getenv("LOG_LEVEL")) // TODO: Pass this default log level through the environment variables through the helm chart

	logrus.Trace("Starting Subscription Service.")

	config := getDefaultArangoDbConfig()
	arango.InitializeArangoDbAdapter(logrus.StandardLogger(), config)

	pubsub.InitializeTopics()
	kafka.StartEventConsumption()

	serverAddress := os.Getenv("APP_SERVER_ADDRESS")

	logger := logrus.WithField("serverAddress", serverAddress)
	logger.Trace("Listening for traffic.")
	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		logger.WithError(err).Panic("Failed to listen for traffic.")
	}

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    30 * time.Second,
			Timeout: 15 * time.Second,
		}),
	)

	signals := helpers.WatchInterruptSignals()
	go func() {
		<-signals
		grpcServer.Stop()
	}()

	jagw.RegisterSubscriptionServiceServer(grpcServer, subscriptionservice.NewServer())
	if err := grpcServer.Serve(lis); err != nil {
		logrus.WithError(err).Panic("Failed to server gRPC server.")
	}
}

func getDefaultArangoDbConfig() arango.ArangoDbConfig {
	return arango.ArangoDbConfig{
		Server:   fmt.Sprintf("http://%s", os.Getenv("ARANGO_ADDRESS")),
		User:     os.Getenv("ARANGO_DB_USER"),
		Password: os.Getenv("ARANGO_DB_PASSWORD"),
		DbName:   os.Getenv("ARANGO_DB_NAME"),
	}
}
