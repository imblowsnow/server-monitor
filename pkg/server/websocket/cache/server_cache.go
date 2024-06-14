package cache

type ServerCache struct {
	cache map[any]interface{}
}

func NewServerCache() *ServerCache {
	return &ServerCache{cache: make(map[any]interface{})}
}

func (c *ServerCache) Set(key any, value interface{}) {
	c.cache[key] = value
}

func (c *ServerCache) Get(key any) interface{} {
	return c.cache[key]
}

func (c *ServerCache) Delete(key any) {
	delete(c.cache, key)
}

func (c *ServerCache) Exists(key any) bool {
	_, ok := c.cache[key]
	return ok
}
