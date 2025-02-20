package cache

import "fmt"

// Manager TODO: CACHE SHARDING?
type Manager struct {
	pool map[string]*Node
}

func NewManager() *Manager {
	return &Manager{
		pool: make(map[string]*Node),
	}
}

func (m *Manager) CacheCreate(name string) string {
	if _, exists := m.pool[name]; exists {
		fmt.Println("Cache already exists:", name)
		return name
	}

	node := (&Node{}).Create()
	m.pool[name] = node
	return name
}

func (m *Manager) CacheDestroy(name string) {
	if _, exists := m.pool[name]; exists {
		m.pool[name].Destroy()
		fmt.Println("Cache destroyed:", name)
	}

	fmt.Println("Cache not found:", name)
}
