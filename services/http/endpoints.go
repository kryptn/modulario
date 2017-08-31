package http

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/kryptn/modulario/proto"
)

func MakeLoginEndpoint(svc HttpService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(proto.LoginRequest)
		resp, err := svc.Login(ctx, req)
		return resp, err
	}
}

func MakeVisitPostEndpoint(svc HttpService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(proto.VisitPostRequest)
		return svc.VisitPost(ctx, req)
	}
}

func MakeViewPostEndpoint(svc HttpService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(proto.ViewPostRequest)
		return svc.ViewPost(ctx, req)
	}
}

func MakeCreatePostEndpoint(svc HttpService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(proto.CreatePostRequest)
		return svc.CreatePost(ctx, req)
	}
}
