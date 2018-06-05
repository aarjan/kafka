package server

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Server describes the hostname & port address of the running server
type Server struct {
	ListenPort uint32
	ListenHost string
}

// Start starts the server with the given configuration
func (s Server) Start(h http.Handler) {
	defer func() {
		if e := recover(); e != nil {
			log.Error(e, "\nServer Failed.")
		}
	}()
	log.WithFields(log.Fields{"Host": s.ListenHost, "Port": s.ListenPort}).Info("Server Started.")

	// run the server
	panic(http.ListenAndServe(s.String(), h))
}

func (s Server) String() string {
	return s.ListenHost + ":" + fmt.Sprintf("%d", s.ListenPort)
}
