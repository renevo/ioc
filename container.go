package ioc

import (
	"reflect"
	"sync"
)

// Container registers and resolves types to a map of named values.
type Container struct {
	mu    sync.RWMutex
	types map[reflect.Type]map[string]reflect.Value
}

// Register will store the supplied type with the value at name.
func (c *Container) Register(name string, t reflect.Type, value any) {
	c.mu.Lock()
	if c.types == nil {
		c.types = make(map[reflect.Type]map[string]reflect.Value)
	}

	if c.types[t] == nil {
		c.types[t] = make(map[string]reflect.Value)
	}

	c.types[t][name] = reflect.ValueOf(value)

	c.mu.Unlock()
}

// Resolve will lookup the type and return a value matching the supplied name.
func (c *Container) Resolve(name string, t reflect.Type) (value any, found bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.types == nil {
		return nil, false
	}
	instanceContainer, found := c.types[t]
	if !found || instanceContainer == nil {
		return nil, false
	}

	if instance, found := instanceContainer[name]; found {
		return instance.Interface(), true
	}
	return nil, false
}

// ResolveAll will lookup the type and return all values registered.
func (c *Container) ResolveAll(t reflect.Type) (values []any) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var instances []any
	if c.types == nil {
		return instances
	}

	instanceContainer, found := c.types[t]
	if !found {
		return instances
	}

	for _, instance := range instanceContainer {
		instances = append(instances, instance.Interface())
	}

	return instances
}
