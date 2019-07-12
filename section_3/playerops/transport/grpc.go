package transport

import (
	"context"

	"github.com/Danr17/microservices_project/section_3/playerops/endpoints"
	"github.com/Danr17/microservices_project/section_3/playerops/pb"
	"github.com/Danr17/microservices_project/section_3/playerops/service"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	createPlayer gt.Handler
	deletePlayer gt.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC PlayerOpsService.
func NewGRPCServer(playerOpsEndpoints endpoints.Endpoints, logger log.Logger) pb.PlayerServiceServer {
	return &gRPCServer{
		createPlayer: gt.NewServer(
			playerOpsEndpoints.CreatePlayerEndpoint,
			decodeCreatePlayerRequest,
			encodeCreatePlayerResponse,
		),
		deletePlayer: gt.NewServer(
			playerOpsEndpoints.DeletePlayerEndpoint,
			decodeDeletePlayerRequest,
			encodeDeletePlayerResponse,
		),
	}
}

func decodeCreatePlayerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreatePlayerRequest)

	player := &service.Player{
		Name:          req.Name.Name,
		Team:          req.Name.Team,
		Nationality:   req.Name.Nationality,
		Position:      req.Name.Position,
		Appearences:   req.Name.Appearences,
		Goals:         req.Name.Goals,
		Assists:       req.Name.Assists,
		Passes:        req.Name.Passes,
		Interceptions: req.Name.Interceptions,
		Tackles:       req.Name.Tackles,
		Fouls:         req.Name.Fouls,
		Price:         req.Name.Price,
	}

	return endpoints.CreatePlayerRequest{PlayerDetail: player}, nil
}

func encodeCreatePlayerResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.CreatePlayerReply)
	return &pb.CreatePlayerReply{Ops: resp.Ops, Err: err2str(resp.Err)}, nil
}

func decodeDeletePlayerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeletePlayerRequest)
	return endpoints.DeletePlayerRequest{Name: req.Name, Team: req.Team}, nil
}

func encodeDeletePlayerResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.DeletePlayerReply)
	return &pb.DeletePlayerReply{Ops: resp.Ops, Err: err2str(resp.Err)}, nil
}

// Helper function is required to translate Go error types to strings,
// which is the type we use in our IDLs to represent errors.

func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func (s *gRPCServer) CreatePlayer(ctx context.Context, req *pb.CreatePlayerRequest) (*pb.CreatePlayerReply, error) {
	_, resp, err := s.createPlayer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CreatePlayerReply), nil
}

func (s *gRPCServer) DeletePLayer(ctx context.Context, req *pb.DeletePlayerRequest) (*pb.DeletePlayerReply, error) {
	_, resp, err := s.deletePlayer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.DeletePlayerReply), nil
}
