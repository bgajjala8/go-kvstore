package cache

import "fmt"

// Manager TODO: CACHE SHARDING?
// This will manage cache in terms of deleting and creating nodes and all api calls
type Manager struct {
	pool map[string]*Vnode
}

func NewManager() *Manager {
	return &Manager{
		pool: make(map[string]*Vnode),
	}
}

func (m *Manager) CacheCreate(name string) string {
	if _, exists := m.pool[name]; exists {
		fmt.Println("Cache already exists:", name)
		return name
	}

	vnode := (&Vnode{}).Create()
	m.pool[name] = vnode
	return name
}

func (m *Manager) CacheDestroy(name string) {
	if _, exists := m.pool[name]; exists {
		m.pool[name].Destroy()
		fmt.Println("Cache destroyed:", name)
	}

	fmt.Println("Cache not found:", name)
}
