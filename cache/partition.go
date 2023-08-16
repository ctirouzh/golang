package cache

import "sync"

type partition struct {
	data map[string]any
	sync.RWMutex
}

func (p *partition) set(key string, value any) {
	p.Lock()
	p.data[key] = value
	p.Unlock()
}

func (p *partition) get(key string) (any, bool) {
	p.RLock()
	defer p.RUnlock()
	v, exist := p.data[key]
	if !exist {
		return nil, false
	}
	return v, true
}
