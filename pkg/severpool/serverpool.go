package serverpool

import (
	"github.com/joshburnsxyz/lb/pkg/backend"
)

type ServerPool struct {
	backends []*backend.Backend
	current uint64
}