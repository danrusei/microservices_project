package service

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/go-kit/kit/log"
	"google.golang.org/api/iterator"
)

var (
	//ErrIterate informs if iteration errors
	ErrIterate = errors.New("can't iterate over the colection documents")

	//ErrExtractDataToStruct informs if unable to extract firestore data to struct
	ErrExtractDataToStruct = errors.New("can't extract the data into a struct with DataTo")
)

//StatsService describe the Stats service
type StatsService interface {
	ListTable(ctx context.Context, league string) ([]*Table, error)
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

func (s *basicService) ListTable(ctx context.Context, league string) ([]*Table, error) {

	var teamTable Table
	var leagueTable []*Table

	leagueDocs := s.dbClient.Collection(league)
	q := leagueDocs.OrderBy("Points", firestore.Desc)
	iter := q.Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, ErrIterate
		}
		if err := doc.DataTo(&teamTable); err != nil {
			return nil, ErrExtractDataToStruct
		}
		leagueTable = append(leagueTable, &teamTable)
	}

	return leagueTable, nil
}

func (s *basicService) ListTeamPlayers(ctx context.Context, teamName string) ([]Player, error) {

	var singlePlayer Player
	var teamPlayers []Player

	teamsDocs := s.dbClient.Collection("Teams")
	q := teamsDocs.Where("team", "array-contains", teamName).OrderBy("player", firestore.Desc)
	iter := q.Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, ErrIterate
		}
		if err := doc.DataTo(&singlePlayer); err != nil {
			return nil, ErrExtractDataToStruct
		}
		teamPlayers = append(teamPlayers, singlePlayer)
	}

	return teamPlayers, nil
}

func (s *basicService) ListPositionPlayers(ctx context.Context, position string) ([]Player, error) {

	var singlePlayer Player
	var teamPlayers []Player

	teamsDocs := s.dbClient.Collection("Teams")
	q := teamsDocs.Where("position", "array-contains", position).OrderBy("team", firestore.Desc)
	iter := q.Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, ErrIterate
		}
		if err := doc.DataTo(&singlePlayer); err != nil {
			return nil, ErrExtractDataToStruct
		}
		teamPlayers = append(teamPlayers, singlePlayer)
	}

	return teamPlayers, nil
}
