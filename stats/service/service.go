package service

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/go-kit/kit/log"
)

//StatsService describe the Stats service
type StatsService interface {
	ListTable(ctx context.Context, league string) ([]Table, error)
	ListTeamPlayers(ctx context.Context, teamName string) ([]Player, error)
	ListPositionPlayers(ctx context.Context, postion string) ([]Player, error)
}

// ** Implementation of the service **

// NewStatsService returns a basic StatsService with all of the expected middlewares wired in.
func NewStatsService(client *firestore.Client, logger log.Logger) StatsService {
	var svc StatsService
	svc = NewBasicService(client)
	svc = LoggingMiddleware(logger)(svc)
	return svc
}

// NewBasicService returns a naive, stateless implementation of StatsService.
func NewBasicService(client *firestore.Client) StatsService {
	return basicService{
		dbClient: client,
	}
}

type basicService struct {
	dbClient *firestore.Client
}

func (s basicService) ListTable(ctx context.Context, league string) ([]Table, error) {

	//implement database request

	return nil, nil
}

func (s basicService) ListTeamPlayers(ctx context.Context, teamName string) ([]Player, error) {

	//implement database request

	return nil, nil
}

func (s basicService) ListPositionPlayers(ctx context.Context, postion string) ([]Player, error) {

	//implement database request

	return nil, nil
}
