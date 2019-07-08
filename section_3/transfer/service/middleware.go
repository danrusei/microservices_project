package service

import (
	"context"

	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(TransferService) TransferService

// LoggingMiddleware takes a logger as a dependency and returns a ServiceMiddleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next TransferService) TransferService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   TransferService
}

func (mw loggingMiddleware) TransferPlayer(ctx context.Context, playerTransfer string, TeamTo string) (ops string, err error) {
	defer func() {
		mw.logger.Log("method", "CreatePlayer", "player", playerTransfer, "err", err)
	}()
	return mw.next.TransferPlayer(ctx, playerTransfer, TeamTo)
}
