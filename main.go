package main

import (
	"flag"
	"github.com/kryptn/modulario/clients/http"
)

var runType = flag.String("ServiceRUnType", "server", "Select server or cli")

func init() {
	flag.Parse()
}

func serveHttp() {
	app := http.Handle
	app()
}

func serveCli() {
	http.Handle()
}

func main() {

	if *runType == "server" {
		serveHttp()
	} else {
		serveCli()
	}

}
