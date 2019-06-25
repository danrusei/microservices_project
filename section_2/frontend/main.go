package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Danr17/microservices_project/section_2/frontend/endpoints"
	"github.com/Danr17/microservices_project/section_2/frontend/service"
	"github.com/Danr17/microservices_project/section_2/frontend/transport"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
)

type frontendServer struct {
	statsSvcAddr string
	statsSvcConn *grpc.ClientConn

	playerSvcAddr string
	playerSvcConn *grpc.ClientConn

	transferSvcAddr string
	transferSvcConn *grpc.ClientConn
}

func main() {

	var httpAddr = ":8080"

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	svc := new(frontendServer)
	mustMapEnv(&svc.statsSvcAddr, "STATS_SERVICE_ADDR")
	//mustMapEnv(&svc.playerSvcAddr, "PLAYER_SERVICE_ADDR")
	//mustMapEnv(&svc.transferSvcAddr, "TRANSFER_SERVICE_ADDR")

	ctx := context.Background()
	mustConnGRPC(ctx, &svc.statsSvcConn, svc.statsSvcAddr)
	//mustConnGRPC(ctx, &svc.playerSvcConn, svc.playerSvcAddr)
	//mustConnGRPC(ctx, &svc.transferSvcConn, svc.transferSvcAddr)

	// Build the layers of the service "onion" from the inside out. First, the
	// business logic service; then, the set of endpoints that wrap the service;
	// and finally, a series of concrete transport adapters

	addservice := service.NewSiteService(logger, svc.statsSvcConn)
	addendpoints := endpoints.MakeSiteEndpoints(addservice)
	httpHandlers := transport.MakeHTTPHandler(addendpoints, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", httpAddr)
		server := &http.Server{
			Addr:    httpAddr,
			Handler: httpHandlers,
		}
		errs <- server.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errs)
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}

func mustConnGRPC(ctx context.Context, conn **grpc.ClientConn, addr string) {
	var err error
	*conn, err = grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Second*3),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
	if err != nil {
		panic(errors.Wrapf(err, "grpc: failed to connect %s", addr))
	}
}
