package cache

type inMemoryScanner struct {
	pair
	pairCh  chan *pair
	closeCh chan struct{}
}

func (s *inMemoryScanner) Close() {
	close(s.closeCh)
}

func (s *inMemoryScanner) Scan() bool {
	p, ok := <-s.pairCh
	if ok {
		s.k, s.v = p.k, p.v
	}
	return ok
}

func (s *inMemoryScanner) Key() string {
	return s.k
}

func (s *inMemoryScanner) Value() []byte {
	return s.v
}

func (c *inMemoryCache) NewScanner() Scanner {
	pairCh := make(chan *pair)
	closeCh := make(chan struct{})
	go func() {
		defer close(pairCh)
		c.mutex.RLock()
		for k, v := range c.c {
			c.mutex.RUnlock()
			select {
			case <-closeCh:
				return
			case pairCh <- &pair{k, v.v}:
			}
			c.mutex.RLock()
		}
		c.mutex.RUnlock()
	}()
	return &inMemoryScanner{pair{}, pairCh, closeCh}
}
