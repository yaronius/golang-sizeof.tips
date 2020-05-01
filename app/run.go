package app

import (
	"flag"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var httpPort int

func init() {
	flag.IntVar(&httpPort, "port", DefaultHttpPort, "Server port")
}

func Run() {
	flag.Parse()

	if err := prepareTemplates(); err != nil {
		log.Error("could not parse html templates, reason -> %s", err.Error())
		return
	}

	bindHttpHandlers()

	log.Infof("Starting server on port %d", httpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil); err != nil {
		log.Errorf("Server error: %s", err)
	}
}
