package transport

import (
	"context"
	"errors"

	"github.com/Danr17/microservices_project/stats/endpoints"
	"github.com/Danr17/microservices_project/stats/pb"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	listTable           gt.Handler
	listTeamPlayers     gt.Handler
	listPositionPlayers gt.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC StatsServiceServer.
func NewGRPCServer(newendpoints endpoints.Endpoints, logger log.Logger) pb.StatsServiceServer {
	return &gRPCServer{
		listTable: gt.NewServer(
			newendpoints.ListTableEndpoint,
			decodeListTableRequest,
			encodeListTableResponse,
		),
		listTeamPlayers: gt.NewServer(
			newendpoints.ListTeamPlayersEndpoint,
			decodeListTeamPlayers,
			encodeListTeamPlayers,
		),
		listPositionPlayers: gt.NewServer(
			newendpoints.ListPositionPlayersEndpoint,
			decodeListPositionPlayers,
			encodeListPositionPlayers,
		),
	}
}

func (s *gRPCServer) ListTable(ctx context.Context, req *pb.TableRequest) (*pb.TableReply, error) {
	_, resp, err := s.listTable.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.TableReply), nil
}

func (s *gRPCServer) ListTeamPlayers(ctx context.Context, req *pb.TeamRequest) (*pb.TeamReply, error) {
	_, resp, err := s.listTeamPlayers.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.TeamReply), nil
}

func (s *gRPCServer) ListPositionPlayers(ctx context.Context, req *pb.PositionRequest) (*pb.PositionReply, error) {
	_, resp, err := s.listPositionPlayers.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.PositionReply), nil
}

func decodeListTableRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.TableRequest)
	return endpoints.TableRequest{League: req.TableName}, nil
}

func encodeListTableResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.TableReply)
	return &pb.TableReply{Teams: resp.Teams, Err: err2str(resp.Err)}, nil
}

func decodeListTeamPlayers(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.TeamRequest)
	return endpoints.TeamRequest{TeamName: req.TeamName}, nil
}

func encodeListTeamPlayers(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.TeamReply)
	return &pb.TeamReply{Players: resp.Players, Err: err2str(resp.Err)}, nil
}

func decodeListPositionPlayers(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.PositionRequest)
	return endpoints.PositionRequest{Position: req.Position}, nil
}

func encodeListPositionPlayers(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.PositionReply)
	return &pb.PositionReply{Players: resp.Players, Err: err2str(resp.Err)}, nil
}

// Helper functions are required to translate Go error types to
// and from strings, which is the type we use in our IDLs to represent errors.

func str2err(s string) error {
	if s == "" {
		return nil
	}
	return errors.New(s)
}

func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
