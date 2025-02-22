package cache

import "fmt"

// Vnode represents a simple in-memory cache virtual node.
// It stores key-value pairs and provides basic cache operations.
type Vnode struct {
	data map[string]string
}

// Create allocates memory for map initialization.
func (v *Vnode) Create() *Vnode {
	v.data = make(map[string]string)
	fmt.Println("Cache created")
	return v
}

// Destroy sets memory block to nil for garbage collection.
func (v *Vnode) Destroy() {
	v.data = nil
}

// Get retrieves the value associated with the given key.
func (v *Vnode) Get(key string) (string, error) {
	if val, exists := v.data[key]; exists {
		return val, nil;
	}
	return "", fmt.Errorf("key does not exist")
}

// Set stores a key-value pair in athe cache.
func (v *Vnode) Set(key string, value string) {
	v.data[key] = value
}

// Delete removes the specified key from the cache.
func (v *Vnode) Delete(key string) {
	delete(v.data,key)
}
