package ioc

import (
	"context"
	"errors"
)

var (
	// ErrContainerNotInContext is used when a container is not found in the context.
	ErrContainerNotInContext = errors.New("container not found in context")
)

type contextKey string

var iocContextKey = contextKey("ioc.Container")

// WithContext will place the supplied container into a context.Context. You can then access the container with ioc.FromContext(ctx).
func WithContext(ctx context.Context, container *Container) context.Context {
	return context.WithValue(ctx, iocContextKey, container)
}

// FromContext will return a container that is stored in the context. If a container is not found, the global default container will be returned.
func FromContext(ctx context.Context) *Container {
	v := ctx.Value(iocContextKey)
	if v == nil {
		return nil
	}

	return v.(*Container)
}

// RegisterToContext will register the supplied value of type T as the default value when using ResolveFromContext().
//
// This function will panic if the context does not contain a container.
func RegisterToContext[T any](ctx context.Context, value T) {
	container := FromContext(ctx)
	if container == nil {
		panic(ErrContainerNotInContext)
	}

	(&GenericContainer[T]{Container: container}).Register(value)
}

// RegisterNamedToContext will register the supplied value of type T with the specified name.
//
// This function will panic if the context does not contain a container.
func RegisterNamedToContext[T any](ctx context.Context, name string, value T) {
	container := FromContext(ctx)
	if container == nil {
		panic(ErrContainerNotInContext)
	}

	(&GenericContainer[T]{Container: container}).RegisterNamed(name, value)
}

// ResolveFromContext will lookup the default registered value.
//
// This function will panic if the context does not contain a container.
func ResolveFromContext[T any](ctx context.Context) (value T, found bool) {
	container := FromContext(ctx)
	if container == nil {
		panic(ErrContainerNotInContext)
	}

	return (&GenericContainer[T]{Container: container}).Resolve()
}

// ResolveNamedFromContext will lookup the value with the specified name.
//
// This function will panic if the context does not contain a container.
func ResolveNamedFromContext[T any](ctx context.Context, name string) (value T, found bool) {
	container := FromContext(ctx)
	if container == nil {
		panic(ErrContainerNotInContext)
	}

	return (&GenericContainer[T]{Container: container}).ResolveNamed(name)
}

// ResolveAllFromContext will lookup and return all values registered.
//
// This function will panic if the context does not contain a container.
func ResolveAllFromContext[T any](ctx context.Context) (values []T) {
	container := FromContext(ctx)
	if container == nil {
		panic(ErrContainerNotInContext)
	}

	return (&GenericContainer[T]{Container: container}).ResolveAll()
}
