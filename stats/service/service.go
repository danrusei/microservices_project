package service

import (
	"context"

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
func NewStatsService(logger log.Logger) StatsService {
	var svc StatsService
	svc = NewBasicService()
	svc = LoggingMiddleware(logger)(svc)
	return svc
}

// NewBasicService returns a naive, stateless implementation of StatsService.
func NewBasicService() StatsService {
	return basicService{}
}

type basicService struct{}

func (s basicService) ListTable(ctx context.Context, league string) ([]Table, error) {

}

func (s basicService) ListTeamPlayers(ctx context.Context, teamName string) ([]Player, error) {

}

func (s basicService) ListPositionPlayers(ctx context.Context, postion string) ([]Player, error) {

}
