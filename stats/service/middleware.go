package service

import (
	"context"

	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(StatsService) StatsService

// LoggingMiddleware takes a logger as a dependency and returns a ServiceMiddleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next StatsService) StatsService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   StatsService
}

func (mw loggingMiddleware) ListTable(ctx context.Context, league string) (t []Table, err error) {
	defer func() {
		mw.logger.Log("method", "Listable", "league", league, "err", err)
	}()
	return mw.next.ListTable(ctx, league)
}

func (mw loggingMiddleware) ListTeamPlayers(ctx context.Context, teamName string) (p []Player, err error) {
	defer func() {
		mw.logger.Log("method", "ListTeamPlayers", "teamName", teamName, "err", err)
	}()
	return mw.next.ListTeamPlayers(ctx, teamName)
}

func (mw loggingMiddleware) ListPositionPlayers(ctx context.Context, position string) (p []Player, err error) {
	defer func() {
		mw.logger.Log("method", "ListPostionPlayers", "position", position, "err", err)
	}()
	return mw.next.ListPositionPlayers(ctx, position)
}
