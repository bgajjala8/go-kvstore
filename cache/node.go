package cache

import "fmt"

// Node CacheNode /** A cache node it what will be created in memory and will throw back point **/
// TODO: ATT TTL STUFF I THINK
// TODO: HOT CACHE COLD CACHE?
type Node struct {
	data map[string]string
}

// Create Returns cache node but who manages cache nodes? in case pointer is lost?
func (c *Node) Create() *Node {
	c.data = make(map[string]string)
	fmt.Println("Cache created")
	return c
}

// Destroy sets cache block to nil
func (c *Node) Destroy() {
	c.data = nil //allows garbage collector to clean up memory
}
