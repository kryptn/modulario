package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	httpSvc "github.com/kryptn/modulario/services/http"
	"time"
)

func BuildHandler() func() {

	r := mux.NewRouter()

	logger := log.Logger{}
	logger.SetOutput(os.Stdout)

	var svc httpSvc.HttpService = httpSvc.MakeHttpService()
	svc = loggingMiddleware{&logger, svc}
	svc = MakeMetricsMiddleware(svc)

	r.Handle(`/visit/{key:[a-zA-Z0-9]{5,12}}`, httptransport.NewServer(
		httpSvc.MakeVisitPostEndpoint(svc),
		decodeVisitEndpoint,
		encodeResponse,
	))

	r.Handle(`/view/{key:[a-zA-Z0-9]{5,12}}`, httptransport.NewServer(
		httpSvc.MakeViewPostEndpoint(svc),
		decodeViewEndpoint,
		encodeResponse,
	))

	r.Handle("/create", httptransport.NewServer(
		httpSvc.MakeCreatePostEndpoint(svc),
		decodeCreateEndpoint,
		encodeResponse,
	)).Methods("POST")

	r.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{
		Handler:      r,
		Addr:         ":5000",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		ErrorLog:     &logger,
	}

	return func() {
		log.Print("Listening on :5000")
		srv.ListenAndServe()
	}
}
