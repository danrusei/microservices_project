package endpoints

import (
	"context"

	"github.com/Danr17/microservices_project/section_3/frontend/service"
	"github.com/go-kit/kit/endpoint"
)

//Endpoints holds all Stats Service enpoints
type Endpoints struct {
	GetTableEndpoint               endpoint.Endpoint
	GetTeamBestPlayersEndpoint     endpoint.Endpoint
	GetPositionBestPlayersEndpoint endpoint.Endpoint
	CreatePlayerEndpoint           endpoint.Endpoint
	DeletePlayerEndpoint           endpoint.Endpoint
	TransferPlayerEndpoint         endpoint.Endpoint
}

//MakeSiteEndpoints initialize all service Endpoints
func MakeSiteEndpoints(s service.SiteService) Endpoints {
	return Endpoints{
		GetTableEndpoint:               makeGetTableEndpoint(s),
		GetTeamBestPlayersEndpoint:     makeGetTeamBestPlayersEndpoint(s),
		GetPositionBestPlayersEndpoint: makeGetPositionBestPlayersEndpoint(s),
		CreatePlayerEndpoint:           makeCreatePlayerEndpoint(s),
		DeletePlayerEndpoint:           makeDeletePlayerEndpoint(s),
		TransferPlayerEndpoint:         makeTransferPlayerEndpoint(s),
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

//CreatePlayerRequest holds the request paramas for CreatePlayer
type CreatePlayerRequest struct {
	NewPlayer service.Player
	TeamName  string
}

//CreatePlayerReply  holds the response paramas for CeeatePLayer
type CreatePlayerReply struct {
	Ops string
	Err error
}

func makeCreatePlayerEndpoint(s service.SiteService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreatePlayerRequest)
		response, err := s.CreatePlayer(ctx, req.NewPlayer, req.TeamName)
		return CreatePlayerReply{Ops: response, Err: err}, nil
	}
}

//DeletePlayerRequest holds the request paramas for DeletePlayer
type DeletePlayerRequest struct {
	DelPlayer string
	TeamName  string
}

//DeletePlayerReply  holds the response paramas for DeletePLayer
type DeletePlayerReply struct {
	Ops string
	Err error
}

func makeDeletePlayerEndpoint(s service.SiteService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeletePlayerRequest)
		response, err := s.DeletePlayer(ctx, req.DelPlayer, req.TeamName)
		return DeletePlayerReply{Ops: response, Err: err}, nil
	}
}

//TransferPlayerRequest holds the request paramas for TransferPlayer
type TransferPlayerRequest struct {
	PlayerName string
	TeamFrom   string
	TeamTo     string
}

//TransferPlayerReply  holds the response paramas for TransferPLayer
type TransferPlayerReply struct {
	Ops string
	Err error
}

func makeTransferPlayerEndpoint(s service.SiteService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(TransferPlayerRequest)
		response, err := s.TransferPlayer(ctx, req.PlayerName, req.TeamFrom, req.TeamTo)
		return TransferPlayerReply{Ops: response, Err: err}, nil
	}
}
