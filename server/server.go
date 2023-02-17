package server

import (
	"fmt"
	"net/http"

	"github.com/joshburnsxyz/lb/serverpool"
)

type Server http.Server

func New(serverPool *serverpool.ServerPool, port int) *Server {
	s := Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(serverPool.Proxy),
	}
	return &s
}