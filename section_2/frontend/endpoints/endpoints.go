package endpoints

import (
	"context"

	"github.com/Danr17/microservices_project/section_2/frontend/service"
	"github.com/go-kit/kit/endpoint"
)

//Endpoints holds all Stats Service enpoints
type Endpoints struct {
	GetTableEndpoint               endpoint.Endpoint
	GetTeamBestPlayersEndpoint     endpoint.Endpoint
	GetPositionBestPlayersEndpoint endpoint.Endpoint
}

//MakeSiteEndpoints initialize all service Endpoints
func MakeSiteEndpoints(s service.SiteService) Endpoints {
	return Endpoints{
		GetTableEndpoint:               makeGetTableEndpoint(s),
		GetTeamBestPlayersEndpoint:     makeGetTeamBestPlayersEndpoint(s),
		GetPositionBestPlayersEndpoint: makeGetPositionBestPlayersEndpoint(s),
	}
}

//TableRequest holds the request params for ListTables
type TableRequest struct {
	League string
}

//TableReply holds the response params for ListTables
type TableReply struct {
	Teams []*service.Table `json:"teams"`
	Err   error            `json:"err"`
}

func makeGetTableEndpoint(s service.SiteService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(TableRequest)
		table, err := s.GetTable(ctx, req.League)
		return TableReply{Teams: table, Err: err}, nil
	}
}

//BestPlayersRequest holds the request paramas for GetTeamBestPlayers
type BestPlayersRequest struct {
	Team string
}

//BestPlayersReply holds the response paramas for GetTeamBestPlayers
type BestPlayersReply struct {
	Players []*service.Player
	Err     error
}

func makeGetTeamBestPlayersEndpoint(s service.SiteService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BestPlayersRequest)
		teamplayers, err := s.GetTeamBestPlayers(ctx, req.Team)
		return BestPlayersReply{Players: teamplayers, Err: err}, nil
	}
}

//BestPositionRequest holds the request paramas for GetBestDefenders
type BestPositionRequest struct {
	Position string `json:"position"`
}

//BestPositionReply  holds the response paramas for GetBestDefenders
type BestPositionReply struct {
	Players []*service.Player
	Err     error
}

func makeGetPositionBestPlayersEndpoint(s service.SiteService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BestPositionRequest)
		defenders, err := s.GetPositionBestPlayers(ctx, req.Position)
		return BestPositionReply{Players: defenders, Err: err}, nil
	}
}
