package endpoints

import (
	"context"

	"github.com/Danr17/microservices_project/section_3/playerops/service"
	"github.com/go-kit/kit/endpoint"
)

//Endpoints holds all PlayerOps Service enpoints
type Endpoints struct {
	CreatePlayerEndpoint endpoint.Endpoint
	DeletePlayerEndpoint endpoint.Endpoint
}

//MakePlayerOpsEndpoints initialize all service Endpoints
func MakePlayerOpsEndpoints(s service.PlayerOpsService) Endpoints {
	return Endpoints{
		CreatePlayerEndpoint: makeCreatePlayerEndpoint(s),
		DeletePlayerEndpoint: makeDeletePlayerEndpoint(s),
	}
}

//CreatePlayerRequest holds the request params for CreatePlayer
type CreatePlayerRequest struct {
	playerDetail []service.Player
}

//CreatePlayerReply holds the response params for CreatePlayer
type CreatePlayerReply struct {
	ops string
	Err error
}

func makeCreatePlayerEndpoint(s service.PlayerOpsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreatePlayerRequest)
		response, err := s.CreatePlayer(ctx, req.playerDetail)
		return CreatePlayerReply{ops: response, Err: err}, nil
	}
}

//DeletePlayerRequest holds the request params for DeletePlayer
type DeletePlayerRequest struct {
	name string
}

//DeletePlayerReply holds the response params for DeletePlayer
type DeletePlayerReply struct {
	ops string
	Err error
}

func makeDeletePlayerEndpoint(s service.PlayerOpsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeletePlayerRequest)
		response, err := s.DeletePlayer(ctx, req.name)
		return DeletePlayerReply{ops: response, Err: err}, nil
	}
}
