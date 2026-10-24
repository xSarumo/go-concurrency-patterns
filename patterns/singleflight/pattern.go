package patterns

import "sync"

type call struct {
	done   chan struct{}
	result any
	err    error
}

type Group struct {
	mu    sync.Mutex
	calls map[string]*call
}

func (g *Group) Do(
	key string,
	fn func() (any, error),
) (any, error) {
	g.mu.Lock()

	if g.calls == nil {
		g.calls = make(map[string]*call)
	}

	if c, exists := g.calls[key]; exists {
		g.mu.Unlock()

		<-c.done

		return c.result, c.err
	}

	c := &call{
		done: make(chan struct{}),
	}

	g.calls[key] = c
	g.mu.Unlock()

	c.result, c.err = fn()

	g.mu.Lock()
	delete(g.calls, key)
	close(c.done)
	g.mu.Unlock()

	return c.result, c.err
}
