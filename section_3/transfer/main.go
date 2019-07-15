package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/firestore"
	"github.com/Danr17/microservices_project/section_3/transfer/endpoints"
	"github.com/Danr17/microservices_project/section_3/transfer/pb"
	"github.com/Danr17/microservices_project/section_3/transfer/service"
	"github.com/Danr17/microservices_project/section_3/transfer/transport"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

func main() {

	var grpcAddr = ":8083"

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	level.Info(logger).Log("msg", "Transfer service started")
	defer level.Info(logger).Log("msg", "Transfer service ended")

	// add database with credentials to run locally
	ctx := context.Background()
	var firestoreClient *firestore.Client
	sa := option.WithCredentialsFile("keys/apps-microservices-68b9b8c44847.json")
	firestoreClient, err := firestore.NewClient(ctx, "apps-microservices", sa)
	if err != nil {
		logger.Log("database", "firestore", "during", "ClientCreation", "err", err)
		os.Exit(1)
	}

	defer firestoreClient.Close()

	// Build the layers of the service "onion" from the inside out. First, the
	// business logic service; then, the set of endpoints that wrap the service;
	// and finally, a series of concrete transport adapters

	addservice := service.NewTransferService(firestoreClient, logger)
	addendpoints := endpoints.MakeTransferEndpoints(addservice)
	grpcServer := transport.NewGRPCServer(addendpoints, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		level.Info(logger).Log("transport", "GRPC", "addr", grpcAddr)
		baseServer := grpc.NewServer()
		pb.RegisterTransferServiceServer(baseServer, grpcServer)
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
