package http

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"runtime"

	"github.com/kryptn/modulario/proto"
	httpSvc "github.com/kryptn/modulario/services/http"
)

type metricsMiddleware struct {
	handlerDuration          *prometheus.SummaryVec
	handlerDurationHistogram *prometheus.HistogramVec

	next httpSvc.HttpService
}

var (
	normDomain = 0.0002
	normMean   = 0.00001
)

func MakeMetricsMiddleware(next httpSvc.HttpService) *metricsMiddleware {

	handlerDuration := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:  "http_svc",
			Subsystem:  "handlers",
			Name:       "handler_duration_seconds",
			Help:       "Handler latency duration",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"handler"},
	)

	handlerDurationHistogram := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "http_svc",
			Subsystem: "handlers",
			Name:      "handler_duration_seconds",
			Help:      "Handler latency duration",
			Buckets:   prometheus.LinearBuckets(normMean-5*normDomain, .5*normDomain, 20),
		},
		[]string{"handler"},
	)

	goroutinesGague := prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace: "http_svc",
			Subsystem: "runtime",
			Name:      "goroutines_count",
			Help:      "Number of goroutines that currently exist.",
		},
		func() float64 { return float64(runtime.NumGoroutine()) },
	)

	prometheus.Register(handlerDuration)
	prometheus.Register(handlerDurationHistogram)
	prometheus.Register(goroutinesGague)

	return &metricsMiddleware{
		handlerDuration,
		handlerDurationHistogram,
		next,
	}
}

func (mw metricsMiddleware) observeDuration(labels prometheus.Labels) func(time.Time) {
	return func(begin time.Time) {
		mw.handlerDuration.With(labels).Observe(time.Since(begin).Seconds())
		mw.handlerDurationHistogram.With(labels).Observe(time.Since(begin).Seconds())
	}
}

func (mw metricsMiddleware) Login(ctx context.Context, req proto.LoginRequest) (proto.LoginResponse, error) {
	defer mw.observeDuration(map[string]string{"handler": "login"})(time.Now())

	return mw.next.Login(ctx, req)
}

func (mw metricsMiddleware) Logout(ctx context.Context) (proto.LogoutResponse, error) {
	defer mw.observeDuration(map[string]string{"handler": "logout"})(time.Now())

	return mw.next.Logout(ctx)
}

func (mw metricsMiddleware) Register(ctx context.Context, req proto.RegisterRequest) (proto.RegisterResponse, error) {
	defer mw.observeDuration(map[string]string{"handler": "register"})(time.Now())

	return mw.next.Register(ctx, req)
}

func (mw metricsMiddleware) VisitPost(ctx context.Context, req proto.VisitPostRequest) (proto.VisitPostResponse, error) {
	defer mw.observeDuration(map[string]string{"handler": "visit_post"})(time.Now())

	return mw.next.VisitPost(ctx, req)
}

func (mw metricsMiddleware) ViewPost(ctx context.Context, req proto.ViewPostRequest) (proto.ViewPostResponse, error) {
	defer mw.observeDuration(map[string]string{"handler": "view_post"})(time.Now())

	return mw.next.ViewPost(ctx, req)
}

func (mw metricsMiddleware) CreatePost(ctx context.Context, req proto.CreatePostRequest) (proto.CreatePostResponse, error) {
	defer mw.observeDuration(map[string]string{"handler": "create_post"})(time.Now())

	return mw.next.CreatePost(ctx, req)
}

func (mw metricsMiddleware) DeletePost(ctx context.Context, req proto.DeletePostRequest) (proto.DeletePostResponse, error) {
	defer mw.observeDuration(map[string]string{"handler": "delete_post"})(time.Now())

	return mw.next.DeletePost(ctx, req)
}
