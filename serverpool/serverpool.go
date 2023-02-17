package serverpool

import (
	"log"
	"net/http"
	"sync/atomic"

	"github.com/joshburnsxyz/lb/backend"
)

type ServerPool struct {
	backends []*backend.Backend
	current  uint64
}

func (s *ServerPool) NextIndex() int {
	return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.backends)))
}

func (s *ServerPool) GetNextPeer() *backend.Backend {
	var nexti = s.NextIndex()
	l := len(s.backends) + nexti
	for i := nexti; i < l; i++ {
		idx := i % len(s.backends)
		if s.backends[idx].IsAlive() {
			if i != nexti {
				atomic.StoreUint64(&s.current, uint64(idx))
			}
			return s.backends[idx]
		}
	}
	return nil
}

func (s *ServerPool) AddBackend(b *backend.Backend) {
	s.backends = append(s.backends, b)
}

func (s *ServerPool) Proxy(w http.ResponseWriter, r *http.Request) {
	peer := s.GetNextPeer()
	if peer != nil {
		peer.ReverseProxy.ServeHTTP(w, r)
	} else {
		http.Error(w, "Service not available", http.StatusServiceUnavailable)
	}
}

func (s *ServerPool) HealthCheck() {
	for _, b := range s.backends {
		status := "up"
		alive := b.Ping()
		b.SetAlive(alive)
		if !alive {
			status = "down"
		}
		log.Printf("%s [%s]\n", b.URL, status)
	}
}

func New() *ServerPool {
	return &ServerPool{}
}
