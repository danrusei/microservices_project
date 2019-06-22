package endpoints

import (
	"context"

	"github.com/Danr17/microservices_project/stats/service"
	"github.com/go-kit/kit/endpoint"
)

//Endpoints holds all Stats Service enpoints
type Endpoints struct {
	ListTableEndpoint           endpoint.Endpoint
	ListTeamPlayersEndpoint     endpoint.Endpoint
	ListPositionPlayersEndpoint endpoint.Endpoint
}

//MakeStatsEndpoints initialize all service Endpoints
func MakeStatsEndpoints(s service.StatsService) Endpoints {
	return Endpoints{
		ListTableEndpoint:           makeListTableEndpoint(s),
		ListTeamPlayersEndpoint:     makeListTeamPLayersEndpoint(s),
		ListPositionPlayersEndpoint: makeListPositionPlayersEnpoint(s),
	}
}

//TableRequest holds the request params for ListTables
type TableRequest struct {
	League string
}

//TableReply holds the response params for ListTables
type TableReply struct {
	Teams []*service.Table
	Err   error
}

func makeListTableEndpoint(s service.StatsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(TableRequest)
		table, err := s.ListTable(ctx, req.League)
		return TableReply{Teams: table, Err: err}, nil
	}
}

//TeamRequest holds the request params for ListTeamPLayers
type TeamRequest struct {
	TeamName string
}

//TeamReply holds the response params for ListTeamPlayers
type TeamReply struct {
	Players []service.Player
	Err     error
}

func makeListTeamPLayersEndpoint(s service.StatsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(TeamRequest)
		teamPlayers, err := s.ListTeamPlayers(ctx, req.TeamName)
		return TeamReply{Players: teamPlayers, Err: err}, nil
	}
}

//PositionRequest holds the request paramas for ListPositionPlayers
type PositionRequest struct {
	Position string
}

//PositionReply holds the response paramas for ListPositionPlayers
type PositionReply struct {
	Players []service.Player
	Err     error
}

func makeListPositionPlayersEnpoint(s service.StatsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PositionRequest)
		positionPlayers, err := s.ListPositionPlayers(ctx, req.Position)
		return PositionReply{Players: positionPlayers, Err: err}, nil
	}

}
