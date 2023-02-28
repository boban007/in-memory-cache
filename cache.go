package cache

import "errors"

type Cache struct {
	values map[string]interface{}
}

func (c *Cache) Set(key string, value interface{}) {
	c.values[key] = value
}

func (c Cache) Get(key string) (interface{}, error) {
	value, exists := c.values[key]
	if exists {
		return value, nil
	}
	return nil, errors.New("key not found")
}

func (c Cache) Delete(key string) error {
	_, exists := c.values[key]
	if exists {
		delete(c.values, key)
		return nil
	}
	return errors.New("key not found")
}

func NewCache() *Cache {
	return &Cache{
		values: make(map[string]interface{}),
	}
}
