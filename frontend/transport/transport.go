package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/Danr17/microservices_project/frontend/endpoints"
)

// NewHTTPHandler wires endpoints to the HTTP transport.
func NewHTTPHandler(siteEndpoints endpoints.Endpoints, logger log.Logger) http.Handler {

	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/table").Handler(kithttp.NewServer(
		siteEndpoints.GetTableEndpoint,
		decodeGetTableRequest,
		encodeGenericResponse,
		options...,
	))

	r.Methods("GET").Path("/bestplayers/{team}").Handler(kithttp.NewServer(
		siteEndpoints.GetTeamBestPlayersEndpoint,
		decodeGetTeamRequest,
		encodeGenericResponse,
		options...,
	))

	r.Methods("GET").Path("/bestdefenders").Handler(kithttp.NewServer(
		siteEndpoints.GetBestDefendersEndpoint,
		decodeGetDefendersRequest,
		encodeGenericResponse,
		options...,
	))

	r.Methods("GET").Path("/bestattackers").Handler(kithttp.NewServer(
		siteEndpoints.GetBestAttackersEndpoint,
		decodeGetAttackersRequest,
		encodeGenericResponse,
		options...,
	))

	r.Methods("GET").Path("/greatpassers").Handler(kithttp.NewServer(
		siteEndpoints.GetGreatPassersEndpoint,
		decodeGetPassersRequest,
		encodeGenericResponse,
		options...,
	))

	return r
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

/// !!! la errorWrapper trebuie sa gasesc erorirle definite de mine in script
