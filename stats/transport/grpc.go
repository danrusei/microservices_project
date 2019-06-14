package transport

import (
	"github.com/Danr17/microservices_project/stats/endpoints"
	"github.com/Danr17/microservices_project/stats/pb"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	listTable           gt.Handler
	listTeamPlayers     gt.Handler
	listPositionPlayers gt.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC StatsServiceServer.
func NewGRPCServer(newendpoints endpoints.Endpoints, logger log.logger) pb.StatsServiceServer {
	return &gRPCServer{
		listTable: gt.NewServer(
			newendpoints.ListTableEndpoint,
			decodeListTableRequest,
			encodeListTableResponse,
		),
		listTeamPlayers: gt.NewServer(
			newendpoints.ListTeamPlayersEndpoint,
			decodeListTeamPLayers
			encodeListTeamPLayers,
		),
		listPositionPlayers: gt.NewServer(
			newendpoints.ListPositionPlayersEndpoint,
			decodeListPositionPlayers,
			encodeListPositionPLayers,
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

//TODO have to add encode and decode functions !!!!!