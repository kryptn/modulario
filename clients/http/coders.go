package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/kryptn/modulario/proto"
)

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func fork(item interface{}, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}
	return item, nil
}

func decodeLoginEndpoint(_ context.Context, r *http.Request) (interface{}, error) {
	var request proto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return fork(request, err)
}

func decodeLogoutEndpoint(_ context.Context, r *http.Request) (interface{}, error) {
	var request proto.LogoutRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return fork(request, err)
}

func decodeRegisterEndpoint(_ context.Context, r *http.Request) (interface{}, error) {
	var request proto.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return fork(request, err)
}

func decodeVisitEndpoint(_ context.Context, r *http.Request) (interface{}, error) {
	var request proto.VisitPostRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return fork(request, err)
}

func decodeViewEndpoint(_ context.Context, r *http.Request) (interface{}, error) {
	var request proto.ViewPostRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return fork(request, err)
}

func decodeCreateEndpoint(_ context.Context, r *http.Request) (interface{}, error) {
	var request proto.CreatePostRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return fork(request, err)

}

func decodeDeleteEndpoint(_ context.Context, r *http.Request) (interface{}, error) {
	var request proto.DeletePostRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	return fork(request, err)
}
