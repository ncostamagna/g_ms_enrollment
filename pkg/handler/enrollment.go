package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/ncostamagna/g_ms_enrollment_ex/internal/enrollment"
	"github.com/ncostamagna/g_ms_enrollment_ex/pkg/response"
)

func NewEnrollmentHTTPServer(ctx context.Context, endpoints enrollment.Endpoints) http.Handler {

	r := mux.NewRouter()

	opts := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Handle("/enrollments", httptransport.NewServer(
		endpoint.Endpoint(endpoints.Create),
		decodeStoreEnrollment,
		encodeResponse,
		opts...,
	)).Methods("POST")

	return r

}

func decodeStoreEnrollment(_ context.Context, r *http.Request) (interface{}, error) {
	var req enrollment.CreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}

	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	r := resp.(response.Response)

	w.WriteHeader(r.StatusCode())
	return json.NewEncoder(w).Encode(resp)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := err.(response.Response)

	w.WriteHeader(resp.StatusCode())

	_ = json.NewEncoder(w).Encode(resp)
}
