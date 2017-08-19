package main

import "github.com/kryptn/modulario/clients/http"

func serveHttp() {
	app := http.Handle
	app()
}

func main() {
	serveHttp()
}