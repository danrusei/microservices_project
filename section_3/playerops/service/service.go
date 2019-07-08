package service

import (
	"context"
	"errors"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/go-kit/kit/log"
)

var (
	//ErrWrite informs if database write error occurs
	ErrWrite = errors.New("can't write the document")

	//ErrDelete informs if unable to delete the player
	ErrDelete = errors.New("can't delete the player from database")
)

//PlayerOpsService  describe the PlayerOps service
type PlayerOpsService interface {
	CreatePlayer(ctx context.Context, playerCreate *Player) (string, error)
	DeletePlayer(ctx context.Context, playerDelete string, teamName string) (string, error)
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

	docName := playerCreate.Team + "_" + strings.Replace(playerCreate.Name, " ", "_", -1)

	wr, err := s.dbClient.Collection("Teams").Doc(docName).Create(ctx, playerCreate)

	if err != nil {
		return "", ErrWrite
	}

	ops := "Write operation completed at " + wr.UpdateTime.String()

	return ops, nil
}

func (s *basicService) DeletePlayer(ctx context.Context, playerDelete string, teamName string) (string, error) {

	docName := teamName + "_" + strings.Replace(playerDelete, " ", "_", -1)

	wr, err := s.dbClient.Collection("Teams").Doc(docName).Delete(ctx)
	if err != nil {
		return "", ErrDelete
	}

	ops := "Deleted " + playerDelete + " at " + wr.UpdateTime.String()

	return ops, nil
}
