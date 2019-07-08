package transport

import (
	"context"

	"github.com/Danr17/microservices_project/section_3/transfer/endpoints"
	"github.com/Danr17/microservices_project/section_3/transfer/pb"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	transferPlayer gt.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC TRansferService.
func NewGRPCServer(transferEndpoints endpoints.Endpoints, logger log.Logger) pb.TransferServiceServer {
	return &gRPCServer{
		transferPlayer: gt.NewServer(
			transferEndpoints.TransferPlayerEndpoint,
			decodeTransferRequest,
			encodeTransferResponse,
		),
	}
}

func decodeTransferRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.PlayerTransferRequest)
	return endpoints.PlayerTransferRequest{PlayerName: req.Name, TeamName: req.Team}, nil
}

func encodeTransferResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.PlayerTransferReply)
	return &pb.PlayerTransferReply{Ops: resp.Ops, Err: err2str(resp.Err)}, nil
}

// Helper function is required to translate Go error types to strings,
// which is the type we use in our IDLs to represent errors.

func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func (s *gRPCServer) TransferPlayer(ctx context.Context, req *pb.PlayerTransferRequest) (*pb.PlayerTransferReply, error) {
	_, resp, err := s.transferPlayer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.PlayerTransferReply), nil
}
