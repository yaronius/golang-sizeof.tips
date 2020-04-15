package log

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// Relative path (from application root) to file
// where application log is stored.
const ApplicationLogFile = "logs/application.log"

// Description of filelog.Writer creation error.
const errCreateLogFile = "failed to create '%s' log file"

// Creates and returns new application logger, ready for use.
func NewApplicationLogger() (*log.Logger, error) {
	lgr := log.New()
	flw, err := os.OpenFile(ApplicationLogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		lgr.SetOutput(flw)
	} else {
		log.Warn("Failed to log to file, using default stderr")
	}
	return lgr, nil
}

// Performs printf() of given pattern with given arguments
// to OS standard error output stream (stderr).
func StdErr(pattern string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, pattern, args...)
}
