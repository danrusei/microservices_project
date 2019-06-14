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

func (s *gRPCServer) ListTable(ctx context.Context, req *pb.TableRequest)
