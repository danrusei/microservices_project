package service

import (
	"context"
	"errors"

	"github.com/Danr17/microservices_project/frontend/pb"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
)

//SiteService describe the Stats service
type SiteService interface {
	GetTable(ctx context.Context, league string) ([]*Table, error)
	GetTeamBestPlayers(ctx context.Context, teamName string) ([]*Player, error)
	GetBestDefenders(ctx context.Context, postion string) ([]*Player, error)
	GetBestAttackers(ctx context.Context, postion string) ([]*Player, error)
	GetGreatPassers(ctx context.Context, postion string) ([]*Player, error)
}

// NewSiteService returns a basic StatsService with all of the expected middlewares wired in.
func NewSiteService(logger log.Logger, conn *grpc.ClientConn) SiteService {
	var svc SiteService
	svc = NewBasicService(conn)
	svc = LoggingMiddleware(logger)(svc)
	return svc
}

// NewBasicService returns a naive, stateless implementation of StatsService.
func NewBasicService(conn *grpc.ClientConn) SiteService {
	return basicService{
		gcStats: pb.NewStatsServiceClient(conn),
	}
}

type basicService struct {
	gcStats pb.StatsServiceClient
}

var (
	//ErrTeamNotFound unable to find the requested team
	ErrTeamNotFound = errors.New("team not found")
	//ErrPLayerNotFound unable to find requested player
	ErrPLayerNotFound = errors.New("player not found")
	//ErrDisplayTable unable to disply table
	ErrDisplayTable = errors.New("unable to display table")
	//ErrDisplayPlayers unable to disply table
	ErrDisplayPlayers = errors.New("unable to display players")
)

//GetTable display final league table
func (s *basicService) GetTable(ctx context.Context, league string) ([]*Table, error) {
	response, err := s.gcStats.ListTable(context.Background(), &pb.TableRequest{
		TableName: league,
	})
	if err != nil {
		return nil, ErrDisplayTable  
	}

	return response, nil
}

//GetTeamBestPLayers diplay top 3 players of a team (one forward, one mid and one defender)
func (s *basicService) GetTeamBestPlayers(ctx context.Context, teamName string) ([]*Player, error) {

	response, err := s.gcStats.ListTeamPlayers(context.Background(), &pb.TeamRequest{
		TeamName: teamName,
	})
	if err != nil {
		return nil, ErrDisplayPlayers 
	}

	return response, nil
}

//GetBestDefenders display top 3 league defenders
func (s *basicService) GetBestDefenders(ctx context.Context, position string) ([]*Player, error) {

	response, err := s.gcStats.ListPositionPlayers(context.Background(), &pb.PositionRequest{
		Position: position,
	})
	if err != nil {
		return nil, ErrDisplayPlayers 
	}

	return response, nil
}

//GetBestAttackers display top 3 league attackers
func (s *basicService) GetBestAttackers(ctx context.Context, position string) ([]*Player, error) {

	response, err := s.gcStats.ListPositionPlayers(context.Background(), &pb.PositionRequest{
		Position: position,
	})
	if err != nil {
		return nil, ErrDisplayPlayers 
	}

	return response, nil
}

//GetGreatPassers display top 3 league passers
func (s *basicService) GetGreatPassers(ctx context.Context, position string) ([]*Player, error) {

	response, err := s.gcStats.ListPositionPlayers(context.Background(), &pb.PositionRequest{
		Position: position,
	})
	if err != nil {
		return nil, ErrDisplayPlayers 
	}

	return response, nil
}
