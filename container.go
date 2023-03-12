package ioc

import (
	"reflect"
	"sync"
)

const _DEFAULT_ = ""

var (
	global = &Container{}
)

type Container struct {
	mu    sync.RWMutex
	types map[reflect.Type]map[string]reflect.Value
}

func (c *Container) Register(name string, t reflect.Type, v any) {
	c.mu.Lock()
	if c.types == nil {
		c.types = make(map[reflect.Type]map[string]reflect.Value)
	}

	if c.types[t] == nil {
		c.types[t] = make(map[string]reflect.Value)
	}

	c.types[t][name] = reflect.ValueOf(v)

	c.mu.Unlock()
}

func (c *Container) Resolve(name string, t reflect.Type) (any, bool) {
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

func (c *Container) ResolveAll(t reflect.Type) []any {
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

func Register[T any](v T) {
	RegisterNamed(_DEFAULT_, v)
}

func RegisterNamed[T any](name string, v T) {
	var instance T
	global.Register(name, reflect.TypeOf(instance), v)
}

func Resolve[T any]() (T, bool) {
	return ResolveNamed[T](_DEFAULT_)
}

func ResolveNamed[T any](name string) (T, bool) {
	var result T

	v, found := global.Resolve(name, reflect.TypeOf(result))
	if found {
		return v.(T), found
	}

	return result, false
}

func ResolveAll[T any]() []T {
	var instance T

	instances := global.ResolveAll(reflect.TypeOf(instance))

	results := make([]T, len(instances))

	for i := 0; i < len(results); i++ {
		results[i] = instances[i].(T)
	}

	return results
}
