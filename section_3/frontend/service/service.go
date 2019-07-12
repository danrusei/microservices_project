package service

import (
	"context"
	"errors"

	"github.com/Danr17/microservices_project/section_3/frontend/pb"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
)

//SiteService describe the Stats service
type SiteService interface {
	GetTable(ctx context.Context, league string) ([]*Table, error)
	GetTeamBestPlayers(ctx context.Context, teamName string) ([]*Player, error)
	GetPositionBestPlayers(ctx context.Context, position string) ([]*Player, error)
	CreatePlayer(ctx context.Context, newplayer *Player) (string, error)
	DeletePlayer(ctx context.Context, delplayer string, teamName string) (string, error)
	TransferPlayer(ctx context.Context, playerName string, teamFrom string, TeamTO string) (string, error)
}

// NewSiteService returns a basic StatsService with all of the expected middlewares wired in.
func NewSiteService(logger log.Logger, conn1 *grpc.ClientConn, conn2 *grpc.ClientConn, conn3 *grpc.ClientConn) SiteService {
	var svc SiteService
	svc = NewBasicService(conn1, conn2, conn3)
	svc = LoggingMiddleware(logger)(svc)
	return svc
}

// NewBasicService returns a naive, stateless implementation of StatsService.
func NewBasicService(conn1 *grpc.ClientConn, conn2 *grpc.ClientConn, conn3 *grpc.ClientConn) SiteService {
	return &basicService{
		gcStats:  pb.NewStatsServiceClient(conn1),
		gcPlayer: pb.NewPlayerServiceClient(conn2),
		gcTrans:  pb.NewTransferServiceClient(conn3),
	}
}

type basicService struct {
	gcStats  pb.StatsServiceClient
	gcPlayer pb.PlayerServiceClient
	gcTrans  pb.TransferServiceClient
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
	//ErrPosition wrong position request
	ErrPosition = errors.New("wrong position, select either Defender, Forward or Midfielder")
)

//GetTable display final league table
func (s *basicService) GetTable(ctx context.Context, league string) ([]*Table, error) {
	resp, err := s.gcStats.ListTable(context.Background(), &pb.TableRequest{
		TableName: league,
	})
	if err != nil {
		return nil, ErrDisplayTable
	}

	teams := make([]*Table, len(resp.Teams))
	for i := range resp.Teams {
		teams[i] = &Table{
			TeamName:    resp.Teams[i].TeamName,
			TeamPlayed:  resp.Teams[i].TeamPlayed,
			TeamWon:     resp.Teams[i].TeamWon,
			TeamDrawn:   resp.Teams[i].TeamDrawn,
			TeamLost:    resp.Teams[i].TeamLost,
			TeamGF:      resp.Teams[i].TeamGF,
			TeamGA:      resp.Teams[i].TeamGA,
			TeamGD:      resp.Teams[i].TeamGD,
			TeamPoints:  resp.Teams[i].TeamPoints,
			TeamCapital: resp.Teams[i].TeamCapital,
		}
	}

	return teams, str2err(resp.Err)
}

//GetTeamBestPLayers diplay top 3 players of a team (one forward, one mid and one defender)
func (s *basicService) GetTeamBestPlayers(ctx context.Context, teamName string) ([]*Player, error) {

	resp, err := s.gcStats.ListTeamPlayers(context.Background(), &pb.TeamRequest{
		TeamName: teamName,
	})
	if err != nil {
		return nil, ErrDisplayPlayers
	}

	players := make([]*Player, len(resp.Players))
	for i := range resp.Players {
		players[i] = makePlayer(resp.Players[i])
	}

	teambestplayers := topteamplayers(players)

	return teambestplayers, str2err(resp.Err)
}

//GetPositionBestPlayers display top 3 league defenders
func (s *basicService) GetPositionBestPlayers(ctx context.Context, position string) ([]*Player, error) {

	resp, err := s.gcStats.ListPositionPlayers(context.Background(), &pb.PositionRequest{
		Position: position,
	})
	if err != nil {
		return nil, err
		//return nil, ErrDisplayPlayers
	}

	players := make([]*Player, len(resp.Players))
	for i := range resp.Players {
		players[i] = makePlayer(resp.Players[i])
	}

	bestpositionplayers, err := topposplayers(players, position)
	if err != nil {
		return nil, ErrPosition
	}

	return bestpositionplayers, str2err(resp.Err)
}

func makePlayer(p *pb.Player) *Player {
	player := &Player{
		Name:          p.Name,
		Team:          p.Team,
		Nationality:   p.Nationality,
		Position:      p.Position,
		Appearences:   p.Appearences,
		Goals:         p.Goals,
		Assists:       p.Assists,
		Passes:        p.Passes,
		Interceptions: p.Interceptions,
		Tackles:       p.Tackles,
		Fouls:         p.Fouls,
	}

	return player
}

func (s *basicService) CreatePlayer(ctx context.Context, newplayer *Player) (string, error) {

	resp, err := s.gcPlayer.CreatePlayer(context.Background(), &pb.CreatePlayerRequest{
		Name: &pb.Player{
			Name:          newplayer.Name,
			Team:          newplayer.Team,
			Nationality:   newplayer.Nationality,
			Position:      newplayer.Position,
			Appearences:   newplayer.Appearences,
			Goals:         newplayer.Goals,
			Assists:       newplayer.Assists,
			Passes:        newplayer.Passes,
			Interceptions: newplayer.Interceptions,
			Tackles:       newplayer.Tackles,
			Fouls:         newplayer.Fouls,
			Price:         newplayer.Price,
		},
	})

	if err != nil {
		return "", err
	}

	return resp.Ops, str2err(resp.Err)
}

func (s *basicService) DeletePlayer(ctx context.Context, delplayer string, teamName string) (string, error) {

	resp, err := s.gcPlayer.DeletePLayer(context.Background(), &pb.DeletePlayerRequest{
		Name: delplayer,
		Team: teamName,
	})

	if err != nil {
		return "", ErrPLayerNotFound
	}

	return resp.Ops, str2err(resp.Err)
}

func (s *basicService) TransferPlayer(ctx context.Context, playerName string, teamFrom string, teamTo string) (string, error) {
	resp, err := s.gcTrans.TransferPlayer(context.Background(), &pb.PlayerTransferRequest{
		Name:     playerName,
		FromTeam: teamFrom,
		ToTeam:   teamTo,
	})

	if err != nil {
		return "", ErrPLayerNotFound
	}

	return resp.Ops, str2err(resp.Err)
}

// Helper function is required to translate Go error types from strings,
// which is the type we use in our IDLs to represent errors.

func str2err(s string) error {
	if s == "" {
		return nil
	}
	return errors.New(s)
}
