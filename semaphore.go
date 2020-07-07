package semaphore

// Semaphore - weightless implementation of Semaphore for multithreading, based on "chan struct{}"
type Semaphore struct {
	ch chan struct{}
}

// NewSemaphore - create a Semaphore
func NewSemaphore(threads int) *Semaphore {
	return &Semaphore{ch: make(chan struct{}, threads)}
}

// Close - close Semaphore
func (s *Semaphore) Close() {
	if s.ch != nil {
		close(s.ch)
		s.ch = nil
	}
}

// Lock - lock @n threads of Semaphore
func (s *Semaphore) Lock(n int) {
	if n < 1 {
		n = 1
	}
	for i := 0; i < n && s.ch != nil; i++ {
		s.ch <- struct{}{}
	}
}

// Unlock - unlock @n threads of Semaphore
func (s *Semaphore) Unlock(n int) {
	if n < 1 {
		n = 1
	}
	for i := 0; i < n && s.ch != nil; i++ {
		<-s.ch
	}
}

// Len - current lenght of Semaphore chan (numbers of active locks)
func (s *Semaphore) Len() int {
	return len(s.ch)
}

// Cap - capacity of Semaphore (total max limit of locks)
func (s *Semaphore) Cap() int {
	return cap(s.ch)
}
