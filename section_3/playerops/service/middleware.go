package service

import (
	"context"

	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(PlayerOpsService) PlayerOpsService

// LoggingMiddleware takes a logger as a dependency and returns a ServiceMiddleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next PlayerOpsService) PlayerOpsService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   PlayerOpsService
}

func (mw loggingMiddleware) CreatePlayer(ctx context.Context, playerCreate *Player) (ops string, err error) {
	defer func() {
		mw.logger.Log("method", "CreatePlayer", "player", playerCreate, "err", err)
	}()
	return mw.next.CreatePlayer(ctx, playerCreate)
}

func (mw loggingMiddleware) DeletePlayer(ctx context.Context, playerDelete string) (ops string, err error) {
	defer func() {
		mw.logger.Log("method", "DeletePlayer", "player", playerDelete, "err", err)
	}()
	return mw.next.DeletePlayer(ctx, playerDelete)
}
