package transport

import (
	"context"
	"errors"

	"github.com/Danr17/microservices_project/section_2/stats/endpoints"
	"github.com/Danr17/microservices_project/section_2/stats/pb"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	listTable           gt.Handler
	listTeamPlayers     gt.Handler
	listPositionPlayers gt.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC StatsServiceServer.
func NewGRPCServer(statsEndpoints endpoints.Endpoints, logger log.Logger) pb.StatsServiceServer {
	return &gRPCServer{
		listTable: gt.NewServer(
			statsEndpoints.ListTableEndpoint,
			decodeListTableRequest,
			encodeListTableResponse,
		),
		listTeamPlayers: gt.NewServer(
			statsEndpoints.ListTeamPlayersEndpoint,
			decodeListTeamPlayers,
			encodeListTeamPlayers,
		),
		listPositionPlayers: gt.NewServer(
			statsEndpoints.ListPositionPlayersEndpoint,
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

	teams := make([]*pb.Table, len(resp.Teams))
	for i := range resp.Teams {
		teams[i] = &pb.Table{
			TeamName:    resp.Teams[i].TeamName,
			TeamPlayed:  resp.Teams[i].TeamPlayed,
			TeamWon:     resp.Teams[i].TeamWon,
			TeamDrawn:   resp.Teams[i].TeamDrawn,
			TeamLost:    resp.Teams[i].TeamLost,
			TeamGF:      resp.Teams[i].TeamGF,
			TeamGA:      resp.Teams[i].TeamGA,
			TeamGD:      resp.Teams[i].TeamGD,
			TeamPoints:  resp.Teams[i].TeamPoints,
			TeamCapital: resp.Teams[i].TeamCapital,
		}
	}

	return &pb.TableReply{Teams: teams, Err: err2str(resp.Err)}, nil
}

func decodeListTeamPlayers(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.TeamRequest)
	return endpoints.TeamRequest{TeamName: req.TeamName}, nil
}

func encodeListTeamPlayers(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.TeamReply)

	players := make([]*pb.Player, len(resp.Players))
	for i := range resp.Players {
		players[i] = &pb.Player{
			Name:          resp.Players[i].Name,
			Team:          resp.Players[i].Team,
			Nationality:   resp.Players[i].Nationality,
			Position:      resp.Players[i].Position,
			Appearences:   resp.Players[i].Appearences,
			Goals:         resp.Players[i].Goals,
			Assists:       resp.Players[i].Assists,
			Passes:        resp.Players[i].Passes,
			Interceptions: resp.Players[i].Interceptions,
			Tackles:       resp.Players[i].Tackles,
			Fouls:         resp.Players[i].Fouls,
		}
	}

	return &pb.TeamReply{Players: players, Err: err2str(resp.Err)}, nil
}

func decodeListPositionPlayers(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.PositionRequest)
	return endpoints.PositionRequest{Position: req.Position}, nil
}

func encodeListPositionPlayers(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.PositionReply)

	players := make([]*pb.Player, len(resp.Players))
	for i := range resp.Players {
		players[i] = &pb.Player{
			Name:          resp.Players[i].Name,
			Team:          resp.Players[i].Team,
			Nationality:   resp.Players[i].Nationality,
			Position:      resp.Players[i].Position,
			Appearences:   resp.Players[i].Appearences,
			Goals:         resp.Players[i].Goals,
			Assists:       resp.Players[i].Assists,
			Passes:        resp.Players[i].Passes,
			Interceptions: resp.Players[i].Interceptions,
			Tackles:       resp.Players[i].Tackles,
			Fouls:         resp.Players[i].Fouls,
		}
	}

	return &pb.PositionReply{Players: players, Err: err2str(resp.Err)}, nil
}

// Helper function is required to translate Go error types to strings,
// which is the type we use in our IDLs to represent errors.

func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
