package endpoints

import (
	"context"

	"github.com/Danr17/microservices_project/section_2/frontend/service"
	"github.com/go-kit/kit/endpoint"
)

//Endpoints holds all Stats Service enpoints
type Endpoints struct {
	GetTableEndpoint           endpoint.Endpoint
	GetTeamBestPlayersEndpoint endpoint.Endpoint
	GetBestDefendersEndpoint   endpoint.Endpoint
	GetBestAttackersEndpoint   endpoint.Endpoint
	GetGreatPassersEndpoint    endpoint.Endpoint
}

//MakeSiteEndpoints initialize all service Endpoints
func MakeSiteEndpoints(s service.SiteService) Endpoints {
	return Endpoints{
		GetTableEndpoint:           makeGetTableEndpoint(s),
		GetTeamBestPlayersEndpoint: makeGetTeamBestPlayersEndpoint(s),
		GetBestDefendersEndpoint:   makeGetBestDefendersEndpoint(s),
		GetBestAttackersEndpoint:   makeGetBestAttackersEndpoint(s),
		GetGreatPassersEndpoint:    makeGetGreatPassersEndpoint(s),
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

//BestPLayersReply holds the response paramas for GetTeamBestPlayers
type BestPLayersReply struct {
	Players []*service.Player
	Err     error
}

func makeGetTeamBestPlayersEndpoint(s service.SiteService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BestPlayersRequest)
		teamplayers, err := s.GetTeamBestPlayers(ctx, req.Team)
		return BestPLayersReply{Players: teamplayers, Err: err}, nil
	}
}

//BestDefendersRequest holds the request paramas for GetBestDefenders
type BestDefendersRequest struct {
	Position string
}

//BestDefendersReply  holds the response paramas for GetBestDefenders
type BestDefendersReply struct {
	Players []*service.Player
	Err     error
}

func makeGetBestDefendersEndpoint(s service.SiteService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BestDefendersRequest)
		defenders, err := s.GetBestDefenders(ctx, req.Position)
		return BestDefendersReply{Players: defenders, Err: err}, nil
	}
}

//BestAttackersRequest holds the request paramas for GetBestAttackers
type BestAttackersRequest struct {
	Position string
}

//BestAttackersReply  holds the response paramas for GetBestAttackers
type BestAttackersReply struct {
	Players []*service.Player
	Err     error
}

func makeGetBestAttackersEndpoint(s service.SiteService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BestAttackersRequest)
		attackers, err := s.GetBestAttackers(ctx, req.Position)
		return BestAttackersReply{Players: attackers, Err: err}, nil
	}
}

//GreatPassersRequest holds the request paramas for GetGreatPassers
type GreatPassersRequest struct {
	Position string
}

//GreatPassersReply holds the response paramas for GetGreatPassers
type GreatPassersReply struct {
	Players []*service.Player
	Err     error
}

func makeGetGreatPassersEndpoint(s service.SiteService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GreatPassersRequest)
		passers, err := s.GetGreatPassers(ctx, req.Position)
		return GreatPassersReply{Players: passers, Err: err}, nil
	}
}
