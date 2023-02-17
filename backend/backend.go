package backend

import (
	"fmt"
	"log"
	"net"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type Backend struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
}

func (b *Backend) IsAlive() (alive bool) {
	b.mux.RLock()
	alive = b.Alive
	b.mux.RUnlock()
	return
}

func (b *Backend) Ping() (alive bool) {
	timeout := 2 * time.Millisecond
	conn, err := net.DialTimeout("tcp", b.URL.Host, timeout)
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		return false
	}
	_ = conn.Close()
	return true
}

func New(url *url.URL) *Backend {
	return &Backend{URL: url}
}
