package http

import (
	"context"
	//"errors"
	//"strings"

	"github.com/kryptn/modulario/proto"
	//"github.com/dgrijalva/jwt-go/request"
	"github.com/davecgh/go-spew/spew"
)

type HttpService interface {
	Login(context.Context, proto.LoginRequest) (proto.LoginResponse, error)
	Logout(ctx context.Context) (proto.LogoutResponse, error)
	Register(context.Context, proto.RegisterRequest) (proto.RegisterResponse, error)
	VisitPost(context.Context, proto.VisitPostRequest) (proto.VisitPostResponse, error)
	ViewPost(context.Context, proto.ViewPostRequest) (proto.ViewPostResponse, error)
	CreatePost(context.Context, proto.CreatePostRequest) (proto.CreatePostResponse, error)
	DeletePost(context.Context, proto.DeletePostRequest) (proto.DeletePostResponse, error)
}

type httpService struct {
	e Engine
}

func MakeHttpService() *httpService {
	e := Engine{Dialect: "sqlite3", Args: "/tmp/gorm.db", LogMode: true}
	e.InitDB()
	e.InitSchema()

	return &httpService{e}
}

func (app *httpService) Login(ctx context.Context, req proto.LoginRequest) (proto.LoginResponse, error) {

	return proto.LoginResponse{}, nil
}

func (app *httpService) Logout(ctx context.Context) (proto.LogoutResponse, error) {

	return proto.LogoutResponse{}, nil
}

func (app *httpService) Register(ctx context.Context, req proto.RegisterRequest) (proto.RegisterResponse, error) {

	return proto.RegisterResponse{}, nil
}

func (app *httpService) VisitPost(ctx context.Context, req proto.VisitPostRequest) (proto.VisitPostResponse, error) {
	resp, err := app.e.VisitKey(req.Key)
	if err != nil {
		return proto.VisitPostResponse{}, err
	}
	return proto.VisitPostResponse{Link: resp.Url}, err
}

func (app *httpService) ViewPost(ctx context.Context, req proto.ViewPostRequest) (proto.ViewPostResponse, error) {
	post, err := app.e.GetPost(req.Key)
	if err != nil {
		return proto.ViewPostResponse{}, err
	}
	spew.Dump(post)
	return proto.ViewPostResponse{Key: post.Key}, nil
}

func (app *httpService) CreatePost(ctx context.Context, req proto.CreatePostRequest) (proto.CreatePostResponse, error) {
	// temp hack until users are a thing
	user := User{}
	user.ID = 0

	post, err := app.e.CreatePost(user, req)
	if err != nil {
		return proto.CreatePostResponse{}, err
	}
	return proto.CreatePostResponse{Key: post.Key}, nil
}

func (app *httpService) DeletePost(ctx context.Context, req proto.DeletePostRequest) (proto.DeletePostResponse, error) {
	return proto.DeletePostResponse{}, app.e.DeletePost(req.Key)
}
