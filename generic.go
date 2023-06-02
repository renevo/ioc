package ioc

import "reflect"

const _DEFAULT_ = ""

// GenericContainer provices some helper functions to register default values as well as provide the reflection options to work with the underlying ioc.Container.
type GenericContainer[T any] struct {
	Container *Container
}

// Register the supplied value of type T as the default value when using Resolve().
func (c *GenericContainer[T]) Register(value T) {
	c.RegisterNamed(_DEFAULT_, value)
}

// RegisterNamed will register the supplied value of type T with the specified name.
func (c *GenericContainer[T]) RegisterNamed(name string, value T) {
	if c.Container == nil {
		panic("GenericContainer.Container is nil")
	}

	var instance T
	c.Container.Register(name, reflect.TypeOf(instance), value)
}

// Resolve will lookup the default registered value.
func (c *GenericContainer[T]) Resolve() (value T, found bool) {
	return c.ResolveNamed(_DEFAULT_)
}

// ResolveNamed will lookup the value with the specified name.
func (c *GenericContainer[T]) ResolveNamed(name string) (value T, found bool) {
	if c.Container == nil {
		panic("GenericContainer.Container is nil")
	}

	var result T
	v, found := c.Container.Resolve(name, reflect.TypeOf(result))
	if found {
		return v.(T), found
	}

	return result, false
}

// ResolveAll will lookup and return all values registered.
func (c *GenericContainer[T]) ResolveAll() (value []T) {
	if c.Container == nil {
		panic("GenericContainer.Container is nil")
	}

	var instance T

	instances := c.Container.ResolveAll(reflect.TypeOf(instance))

	results := make([]T, len(instances))

	for i := 0; i < len(results); i++ {
		results[i] = instances[i].(T)
	}

	return results
}
