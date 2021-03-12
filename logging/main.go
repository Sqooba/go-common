package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
	"time"
)

// Wrapper around https://github.com/sirupsen/logrus
// including facilities such as runtime logging level modification,
// http handler to change verbosity as well as simple client using the handler.

const (
	// Ok is the "ok" ServiceCheck status
	Ok float64 = 0
	// Warn is the "warning" ServiceCheck status
	Warn float64 = 1
	// HTTP path to expose health check
	DefaultLogVerbosityPath = "/debug/verbosity"
	Localhost               = "127.0.0.1"
)

// NewLogger returns logrus standard logger.
func NewLogger() *logrus.Logger {
	return logrus.StandardLogger()
}

// LogLevelHandler allow to change the log level at runtime using http PUT request:
// curl -X GET http://... returns the current log level
// curl -X PUT -d debug http://... set logging level to debug.
func LogLevelHandler(log *logrus.Logger) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			http.Error(w,
				fmt.Sprintf("%s\n", log.Level), http.StatusOK)
		} else if req.Method == http.MethodPut {
			// A request must have a body.
			if req.Body == nil {
				http.Error(w, "Ignoring request. Required non-empty request body.\n", http.StatusBadRequest)
				return
			}
			defer req.Body.Close()
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				http.Error(w, fmt.Sprintf("Got an error reading body, %v\n", err), http.StatusBadRequest)
				return
			}

			newLevel := string(body)
			err = SetLogLevel(log, newLevel)

			if err != nil {
				http.Error(w, fmt.Sprintf("Got an error parsing body, %v\n", err), http.StatusBadRequest)
				return
			}
			http.Error(w, fmt.Sprintf("New log level %s set\n", newLevel), http.StatusOK)
			return
		} else {
			http.Error(w, fmt.Sprintf("Ignoring request. Required method is \"GET\" or \"PUT\", but got \"%s\".\n", req.Method), http.StatusMethodNotAllowed)
			return
		}
	}
}

// InitVerbosityHandler is a convenience function to set LogLevelHandler to given http server
func InitVerbosityHandler(log *logrus.Logger, mux *http.ServeMux) {
	mux.HandleFunc(DefaultLogVerbosityPath, LogLevelHandler(log))
}

// SetLogLevel changes the log level of the given logger.
func SetLogLevel(log *logrus.Logger, level string) error {
	l, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	atomic.StoreUint32((*uint32)(&log.Level), uint32(l))
	return nil
}

// IsDebugEnabled returns true if the log level if debug or more verbose (i.e. trace)
func IsDebugEnabled(log *logrus.Logger) bool {
	return log.Level >= logrus.DebugLevel
}

// SetRemoteLogLevelAndExit is a helper function which calls LogLevelHandler
// to update the verbosity of a running process.
func SetRemoteLogLevelAndExit(log *logrus.Logger, port string, logLevel string) {

	l, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Errorf("Got an error while parsing log level to %s. Err = %v", logLevel, err)
		os.Exit(int(Warn))
	}

	request, err := http.NewRequest(http.MethodPut,
		fmt.Sprintf("http://%s:%s%s", Localhost, port, DefaultLogVerbosityPath),
		strings.NewReader(l.String()))
	if err != nil {
		log.Errorf("Got an error while creating the http PUT request. Err = %v", err)
		os.Exit(int(Warn))
	}
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Errorf("Got an error while calling the http PUT request. Err = %v", err)
		os.Exit(int(Warn))
	}

	if resp.StatusCode == http.StatusOK {
		log.Debugf("Properly set log level to %s", logLevel)
		os.Exit(int(Ok))
	} else {
		log.Warnf("Got an unexpected status code %v. The verbosity was most likely not changed", resp.StatusCode)
		os.Exit(int(Warn))
	}
}
