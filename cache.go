package main

import "fmt"

type Cache interface {
	cacheCreate() *CacheNode
	cacheDestroy()
}

// CacheNode /** A cache node it what will be created in memory and will throw back point **/
type CacheNode struct {
	name string
	data map[string]string
}

// Returns cache node but who manages cache nodes? in case pointer is lost?
func (c *CacheNode) cacheCreate() *CacheNode {
	c.data = make(map[string]string)
	fmt.Println("Cache created:", c.name)
	return c
}

// Destroy sets cache block to nil
func (c *CacheNode) cacheDestroy() {
	c.data = nil //allows garbage collector to clean up memory
}
