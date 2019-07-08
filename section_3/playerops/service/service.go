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

//PlayerOpsService  describe the PlayerOps service
type PlayerOpsService interface {
	CreatePlayer(ctx context.Context, playerCreate *Player) (string, error)
	DeletePlayer(ctx context.Context, playerDelete string) (string, error)
}

// ** Implementation of the service **

// NewPlayerOpsService returns a basic PLayerOpsService with all of the expected middlewares wired in.
func NewPlayerOpsService(client *firestore.Client, logger log.Logger) PlayerOpsService {
	var svc PlayerOpsService
	svc = NewBasicService(client)
	svc = LoggingMiddleware(logger)(svc)
	return svc
}

// NewBasicService returns a naive, stateless implementation of PlayerOpsService.
func NewBasicService(client *firestore.Client) PlayerOpsService {
	return &basicService{
		dbClient: client,
	}
}

type basicService struct {
	dbClient *firestore.Client
}

func (s *basicService) CreatePlayer(ctx context.Context, playerCreate *Player) (string, error) {

	return "", nil
}

func (s *basicService) DeletePlayer(ctx context.Context, playerDelete string) (string, error) {

	return "", nil
}
