package service

import (
	"context"

	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(SiteService) SiteService

// LoggingMiddleware takes a logger as a dependency and returns a ServiceMiddleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next SiteService) SiteService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   SiteService
}

func (mw loggingMiddleware) GetTable(ctx context.Context, league string) (t []*Table, err error) {
	defer func() {
		mw.logger.Log("method", "GetTable", "league", league, "err", err)
	}()
	return mw.next.GetTable(ctx, league)
}

func (mw loggingMiddleware) GetTeamBestPlayers(ctx context.Context, teamName string) (p []*Player, err error) {
	defer func() {
		mw.logger.Log("method", "GetTeamBestPlayers", "teamName", teamName, "err", err)
	}()
	return mw.next.GetTeamBestPlayers(ctx, teamName)
}

func (mw loggingMiddleware) GetPositionBestPlayers(ctx context.Context, position string) (p []*Player, err error) {
	defer func() {
		mw.logger.Log("method", "GetPositionBestPlayers", "position", position, "err", err)
	}()
	return mw.next.GetPositionBestPlayers(ctx, position)
}
