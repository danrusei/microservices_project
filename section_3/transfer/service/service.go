package service

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/go-kit/kit/log"
)

var (
	//ErrIterate informs if iteration errors
	ErrIterate = errors.New("can't iterate over the colection documents")

	//ErrExtractDataToStruct informs if unable to extract firestore data to struct
	ErrExtractDataToStruct = errors.New("can't extract the data into a struct with DataTo")
)

//TransferService  describe the Transfer service
type TransferService interface {
	TransferPlayer(ctx context.Context, playerTransfer string, TeamTo string) (string, error)
}

// ** Implementation of the service **

// NewTransferService returns a basic TransferService with all of the expected middlewares wired in.
func NewTransferService(client *firestore.Client, logger log.Logger) TransferService {
	var svc TransferService
	svc = NewBasicService(client)
	svc = LoggingMiddleware(logger)(svc)
	return svc
}

// NewBasicService returns a naive, stateless implementation of PlayerOpsService.
func NewBasicService(client *firestore.Client) TransferService {
	return &basicService{
		dbClient: client,
	}
}

type basicService struct {
	dbClient *firestore.Client
}

func (s *basicService) TransferPlayer(ctx context.Context, playerTransfer string, TeamTo string) (string, error) {

	return "", nil
}
