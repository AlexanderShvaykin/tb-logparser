package sync_buff

import (
	"bytes"
	"sync"
)

type Buff struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (s *Buff) Write(p []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.buf.Write(p)
}

func (s *Buff) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.buf.Reset()
}

func (s *Buff) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.buf.String()
}
