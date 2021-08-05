package main

import (
	"context"
	"io"
	"log"
	"net"
	"os"

	psproto "gitlab.ost.ch/ins/jalapeno-api/push-service/proto/push-service"
	tsproto "gitlab.ost.ch/ins/jalapeno-api/push-service/proto/tsdb-feeder"
	"google.golang.org/grpc"
)

type pushServiceServer struct {
	psproto.UnimplementedPushServiceServer
}

func newServer() *pushServiceServer {
	s := &pushServiceServer{}
	return s
}

func main() {
	log.Print("Starting Push Service ...")
	//Start gRPC server for SR-Apps
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
	grpcServer := grpc.NewServer()
	psproto.RegisterPushServiceServer(grpcServer, newServer())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}

func (s *pushServiceServer) SubscribeToDataRates(ipv4Addresses *psproto.IPv4Addresses, responseStream psproto.PushService_SubscribeToDataRatesServer) error {
	log.Printf("SR-App subscribing to DataRates\n")

	//Call SubscribeToDataRate on TSDBFeeder
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(os.Getenv("TSDB_FEEDER_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	client := tsproto.NewTsdbFeederClient(conn)
	message := &tsproto.IPv4Addresses{Ipv4Address: ipv4Addresses.Ipv4Address}
	stream, err := client.SubscribeToDataRates(context.Background(), message)

	if err != nil {
		log.Fatalf("Error when calling SubscribeToDataRate on TSDBFeeder: %s", err)
	}

	for {
		dataRate, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetNodes(_) = _, %v", client, err)
		}
		responseStream.Send(&psproto.DataRate{DataRate: dataRate.DataRate, Ipv4Address: dataRate.Ipv4Address})
	}
	log.Printf("Subscription to DataRates ended")
	return nil
}