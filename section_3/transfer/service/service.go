package service

import (
	"context"
	"errors"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/go-kit/kit/log"
)

var (
	//ErrIterate informs if iteration errors
	ErrIterate = errors.New("can't iterate over the colection documents")

	//ErrExtractDataToStruct informs if unable to extract firestore data to struct
	ErrExtractDataToStruct = errors.New("can't extract the data into a struct with DataTo")

	//ErrWrite informs if database write error occurs
	ErrWrite = errors.New("can't write the document")

	//ErrDelete informs if unable to delete the player
	ErrDelete = errors.New("can't delete the player from database")
)

//TransferService  describe the Transfer service
type TransferService interface {
	TransferPlayer(ctx context.Context, playerTransfer string, TeamFrom string, TeamTo string) (string, error)
}

// ** Implementation of the service **

// NewTransferService returns a basic TransferService with all of the expected middlewares wired in.
func NewTransferService(client *firestore.Client, logger log.Logger) TransferService {
	var svc TransferService
	svc = NewBasicService(client)
	svc = LoggingMiddleware(logger)(svc)
	return svc
}

// NewBasicService returns a naive, stateless implementation of PlayerOpsService.
func NewBasicService(client *firestore.Client) TransferService {
	return &basicService{
		dbClient: client,
	}
}

type basicService struct {
	dbClient *firestore.Client
}

func (s *basicService) TransferPlayer(ctx context.Context, playerTransfer string, TeamFrom string, TeamTo string) (string, error) {

	var teamToTable Table
	var teamFromTable Table
	var playerStruct Player

	docDelete := TeamFrom + "_" + strings.Replace(playerTransfer, " ", "_", -1)
	docCreate := TeamTo + "_" + strings.Replace(playerTransfer, " ", "_", -1)

	//check player price
	//check teamTO Capital
	//if price <= Capital then transfer is allowed
	//delete TeamFrom_player; create TeamTo_player(change Team field)
	//addjust teams money if transfer is done

	teamTOtransfer, err := s.dbClient.Collection("League").Doc(TeamTo).Get(ctx)
	if err != nil {
		return "", err
	}
	if err := teamTOtransfer.DataTo(&teamToTable); err != nil {
		return "", ErrExtractDataToStruct
	}

	teamFROMtransfer, err := s.dbClient.Collection("League").Doc(TeamFrom).Get(ctx)
	if err != nil {
		return "", err
	}
	if err := teamFROMtransfer.DataTo(&teamFromTable); err != nil {
		return "", ErrExtractDataToStruct
	}

	playerTOtransfer, err := s.dbClient.Collection("Teams").Doc(docDelete).Get(ctx)
	if err != nil {
		return "", err
	}
	if err := playerTOtransfer.DataTo(&playerStruct); err != nil {
		return "", ErrExtractDataToStruct
	}

	if teamToTable.TeamCapital < playerStruct.Price {
		return "The team can't afford the player", nil
	}

	_, err = s.dbClient.Collection("Teams").Doc(docDelete).Delete(ctx)
	if err != nil {
		return "", ErrDelete
	}

	_, err = s.dbClient.Collection("Teams").Doc(docCreate).Create(ctx, playerStruct)

	if err != nil {
		return "", ErrWrite
	}

	toTeamCapital := int(teamToTable.TeamCapital - playerStruct.Price)
	_, err = s.dbClient.Collection("League").Doc(TeamTo).Set(ctx, map[string]interface{}{
		"Capital": toTeamCapital}, firestore.MergeAll)
	if err != nil {
		return "", err
	}

	fromTeamCapital := int(teamToTable.TeamCapital + playerStruct.Price)
	wr, err := s.dbClient.Collection("League").Doc(TeamFrom).Set(ctx, map[string]interface{}{
		"Capital": fromTeamCapital}, firestore.MergeAll)
	if err != nil {
		return "", err
	}

	ops := "Player " + playerStruct.Name + " has been transfered to " + teamToTable.TeamName + " from " + teamFromTable.TeamName + " at " + wr.UpdateTime.String()

	return ops, nil
}
