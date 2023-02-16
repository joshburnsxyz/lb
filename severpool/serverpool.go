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
