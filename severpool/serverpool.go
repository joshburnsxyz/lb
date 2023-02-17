package serverpool

import (
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

func New() (*ServerPool) {
	return &ServerPool{}
}