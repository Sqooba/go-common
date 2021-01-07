package healthchecks

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"syscall"

	"github.com/docker/distribution/health"
)

// Wrapper around https://github.com/docker/distribution/health
// including facilities such as simple client usable as docker healthcheck command

const (
	// Ok is the "ok" ServiceCheck status
	Ok float64 = 0
	// Warn is the "warning" ServiceCheck status
	Warn float64 = 1
	// Critical is the "critical" ServiceCheck status
	Critical float64 = 2
	// Unknown is the "unknown" ServiceCheck status
	Unknown float64 = 3
	// HTTP path to expose health check
	HealthCheckPath        = "/debug/health"
	Localhost       string = "127.0.0.1"
)

// RunHealthCheckAndExit is a facility function to not add another binary file in the container
// which is capable of calling the healthcheck endpoint.
func RunHealthCheckAndExit(port string) {
	response, err := http.Get(fmt.Sprintf("http://%s:%s%s", Localhost, port, HealthCheckPath))
	if err != nil || response.StatusCode != http.StatusOK {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		log.Printf("Got an error calling http://%s:%s%s, status=%v, err=%s, body=%s",
			Localhost, port, HealthCheckPath, response.StatusCode, err, string(body))
		os.Exit(int(Warn))
	}
	os.Exit(int(Ok))
}

// AlwaysOkHealthcheckFuncHandler returns always OK (200) status code
func AlwaysOkHealthcheckFuncHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var writeErr error
		w.WriteHeader(http.StatusOK)
		_, writeErr = w.Write([]byte{})

		if writeErr != nil {
			log.Printf("Could not write response: %v", writeErr)
		}
	})
}

// RunHealthcheckHTTPServer starts HTTP server at given port
// and handle /debug/health endpoint, with the default
// AlwaysOkHealthcheckFuncHandler handler.
func RunHealthcheckHTTPServer(healthCheckPort string) *http.ServeMux {
	return RunHealthcheckHTTPServerWithHandler(healthCheckPort, AlwaysOkHealthcheckFuncHandler())
}

// RunHealthcheckHTTPServerWithHandler starts HTTP server at given port
// and handle /debug/health endpoint with the given handler,
// DockerHealthStatusHandler for instance.
func RunHealthcheckHTTPServerWithHandler(healthCheckPort string, healthCheckHAndler http.Handler) *http.ServeMux {

	mux := http.NewServeMux()

	mux.Handle(HealthCheckPath, healthCheckHAndler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", healthCheckPort),
		Handler: mux,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	return mux
}

// DockerHealthStatusHandler wraps the default docker StatusHandler function
// in a http.Handler to be usable here.
func DockerHealthStatusHandler() http.Handler {
	return http.HandlerFunc(health.StatusHandler)
}

// RegisterHealthCheck registers a health function
func RegisterHealthCheck(name string, checkF func() error) {
	health.RegisterFunc(name, checkF)
}

// SendSigIntWhenError returns a function which can send a sigint
func SendSigIntWhenError(checkF func() error) func() error {
	return func() error {
		err := checkF()
		if err != nil {
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		}
		return err
	}
}
