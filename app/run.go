package app

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func Run() {
	appLog = log.New()

	if err := prepareTemplates(); err != nil {
		log.Error("could not parse html templates, reason -> %s", err.Error())
		return
	}

	httpPort := os.Getenv("_GO_HTTP")
	if httpPort == "" {
		httpPort = DefaultHttpPort
	}

	bindHttpHandlers()

	appLog.Infof("Starting server on port %s...", httpPort)
	if err := http.ListenAndServe(httpPort, nil); err != nil {
		log.Error(err)
	}
}
