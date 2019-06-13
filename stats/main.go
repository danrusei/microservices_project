package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"

	"github.com/Danr17/microservices_project/stats/pb"
	"github.com/Danr17/microservices_project/stats/service"
)

func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	// Build the layers of the service "onion" from the inside out. First, the
	// business logic service; then, the set of endpoints that wrap the service;
	// and finally, a series of concrete transport adapters

	service := service.NewStatsService(logger)
	endpoints := enpoints.MakeStatsEndpoints(service)
	grpcServer := transport.NewGRPCServer(endpoints, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		level.Info(logger).Log("transport", "GRPC", "addr", *grpcAddr)
		baseServer := grpc.NewServer()
		pb.RegisterStatsServiceServer(baseServer, grpcServer)
	}()

	level.Error(logger).Log("exit", <-errs)

}
