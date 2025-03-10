package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVnode_Create(t *testing.T) {
	v := &Vnode{}
	v.Create()

	assert.NotNil(t, v.data, "Cache data map should not be nil")
}

func TestVnode_Destroy(t *testing.T) {
	v := setupCache()
	v.Destroy()

	assert.Nil(t, v.data, "Cache instance should be nil")
}

func setupCache() *Vnode {
	v := &Vnode{}
	v.Create()
	return v
}

func destroyCache(v *Vnode) {
	v.Destroy()
}

func TestVnode_Set(t *testing.T) {
	v := setupCache()

	v.Set("foo", "bar")
	value, exists := v.data["foo"]

	assert.True(t, exists, "Key 'foo' does not exist in the cache")
	assert.Equal(t, "bar", value, "Set keys don't match keys in map")

	destroyCache(v)
}

func TestVnode_Get(t *testing.T) {
	v := setupCache()

	v.data["foo"] = "bar"
	value, error := v.Get("foo")

	assert.Nil(t, error, "Get 'foo' should not return an error")
	assert.Equal(t, "bar", value, "Set keys don't match keys in map")

	destroyCache(v)
}

func TestVnode_Delete(t *testing.T) {
	v := setupCache()

	v.data["foo"] = "bar"
	v.Delete("foo")

	value, exists := v.data["foo"]

	assert.Equal(t, "", value, "Value for key 'foo' should be nil")
	assert.True(t, !exists, "Key 'foo' should not exist in the cache")

	destroyCache(v)
}
