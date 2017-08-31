package http

import (
	"context"
	"log"
	"time"

	"github.com/kryptn/modulario/proto"
	httpSvc "github.com/kryptn/modulario/services/http"
)

type loggingMiddleware struct {
	logger *log.Logger
	next   httpSvc.HttpService
}

func (mw loggingMiddleware) Login(ctx context.Context, req proto.LoginRequest) (proto.LoginResponse, error) {
	defer func(begin time.Time) {
		mw.logger.Printf("Login called, took %s", time.Since(begin))
	}(time.Now())

	return mw.next.Login(ctx, req)
}

func (mw loggingMiddleware) Logout(ctx context.Context) (proto.LogoutResponse, error) {
	defer func(begin time.Time) {
		mw.logger.Printf("Logout called, took %s", time.Since(begin))
	}(time.Now())

	return mw.next.Logout(ctx)
}

func (mw loggingMiddleware) Register(ctx context.Context, req proto.RegisterRequest) (proto.RegisterResponse, error) {
	defer func(begin time.Time) {
		mw.logger.Printf("Register called, took %s", time.Since(begin))
	}(time.Now())

	return mw.next.Register(ctx, req)
}

func (mw loggingMiddleware) VisitPost(ctx context.Context, req proto.VisitPostRequest) (proto.VisitPostResponse, error) {
	defer func(begin time.Time) {
		mw.logger.Printf("VisitPost called, took %s", time.Since(begin))
	}(time.Now())

	return mw.next.VisitPost(ctx, req)
}

func (mw loggingMiddleware) ViewPost(ctx context.Context, req proto.ViewPostRequest) (proto.ViewPostResponse, error) {
	defer func(begin time.Time) {
		mw.logger.Printf("ViewPost called, took %s", time.Since(begin))
	}(time.Now())

	return mw.next.ViewPost(ctx, req)
}

func (mw loggingMiddleware) CreatePost(ctx context.Context, req proto.CreatePostRequest) (proto.CreatePostResponse, error) {
	defer func(begin time.Time) {
		mw.logger.Printf("CreatePost called, took %s", time.Since(begin))
	}(time.Now())

	return mw.next.CreatePost(ctx, req)
}

func (mw loggingMiddleware) DeletePost(ctx context.Context, req proto.DeletePostRequest) (proto.DeletePostResponse, error) {
	defer func(begin time.Time) {
		mw.logger.Printf("Register called, took %s", time.Since(begin))
	}(time.Now())

	return mw.next.DeletePost(ctx, req)
}
