package ioc

import "context"

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
		return global
	}

	return v.(*Container)
}

// RegisterToContext will register the supplied value of type T as the default value when using ResolveFromContext().
func RegisterToContext[T any](ctx context.Context, value T) {
	(&GenericContainer[T]{Container: FromContext(ctx)}).Register(value)
}

// RegisterNamedToContext will register the supplied value of type T with the specified name.
func RegisterNamedToContext[T any](ctx context.Context, name string, value T) {
	(&GenericContainer[T]{Container: FromContext(ctx)}).RegisterNamed(name, value)
}

// ResolveFromContext will lookup the default registered value.
func ResolveFromContext[T any](ctx context.Context) (value T, found bool) {
	return (&GenericContainer[T]{Container: FromContext(ctx)}).Resolve()
}

// ResolveNamedFromContext will lookup the value with the specified name.
func ResolveNamedFromContext[T any](ctx context.Context, name string) (value T, found bool) {
	return (&GenericContainer[T]{Container: FromContext(ctx)}).ResolveNamed(name)
}

// ResolveAllFromContext will lookup and return all values registered.
func ResolveAllFromContext[T any](ctx context.Context) (values []T) {
	return (&GenericContainer[T]{Container: FromContext(ctx)}).ResolveAll()
}
