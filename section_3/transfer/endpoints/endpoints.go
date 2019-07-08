package endpoints

import (
	"context"

	"github.com/Danr17/microservices_project/section_3/transfer/service"
	"github.com/go-kit/kit/endpoint"
)

//Endpoints holds all PlayerOps Service enpoints
type Endpoints struct {
	TransferPlayerEndpoint endpoint.Endpoint
}

//MakeTransferEndpoints initialize all service Endpoints
func MakeTransferEndpoints(s service.TransferService) Endpoints {
	return Endpoints{
		TransferPlayerEndpoint: makeTransferPlayerEndpoint(s),
	}
}

//PlayerTransferRequest holds the request params for PlayerTransfer
type PlayerTransferRequest struct {
	PlayerName string
	FromTeam   string
	ToTeam     string
}

//PlayerTransferReply holds the response params for PlayerTransfer
type PlayerTransferReply struct {
	Ops string
	Err error
}

func makeTransferPlayerEndpoint(s service.TransferService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PlayerTransferRequest)
		response, err := s.TransferPlayer(ctx, req.PlayerName, req.FromTeam, req.ToTeam)
		return PlayerTransferReply{Ops: response, Err: err}, nil
	}
}
