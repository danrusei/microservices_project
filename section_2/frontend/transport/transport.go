package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/Danr17/microservices_project/section_2/frontend/endpoints"
	"github.com/Danr17/microservices_project/section_2/frontend/service"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler")
)

// MakeHTTPHandler wires endpoints to the HTTP transport.
func MakeHTTPHandler(siteEndpoints endpoints.Endpoints, logger log.Logger) http.Handler {

	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/table").Handler(kithttp.NewServer(
		siteEndpoints.GetTableEndpoint,
		decodeGetTableRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/bestplayers/{team}").Handler(kithttp.NewServer(
		siteEndpoints.GetTeamBestPlayersEndpoint,
		decodeGetTeamRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/bestposition").Handler(kithttp.NewServer(
		siteEndpoints.GetPositionBestPlayersEndpoint,
		decodeGetPositionRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeGetTableRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoints.TableRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetTeamRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	team, ok := vars["team"]
	if !ok {
		return nil, ErrBadRouting
	}
	return endpoints.BestPlayersRequest{Team: team}, nil
}

func decodeGetPositionRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoints.BestPositionRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error. For more information, read the
// big comment in endpoints.go.
type errorer interface {
	error() error
}

// encodeResponse is the common method to encode all response types to the
// client. Since we're using JSON, there's no reason to provide anything more specific.
// It's certainly possible to specialize on a per-response (per-method) basis.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case service.ErrTeamNotFound, service.ErrPLayerNotFound:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
