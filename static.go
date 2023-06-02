package ioc

var global = &Container{}

// Register the supplied value of type T as the default value when using Resolve().
func Register[T any](value T) {
	(&GenericContainer[T]{Container: global}).Register(value)
}

// RegisterNamed will register the supplied value of type T with the specified name.
func RegisterNamed[T any](name string, value T) {
	(&GenericContainer[T]{Container: global}).RegisterNamed(name, value)
}

// Resolve will lookup the default registered value.
func Resolve[T any]() (value T, found bool) {
	return (&GenericContainer[T]{Container: global}).Resolve()
}

// ResolveNamed will lookup the value with the specified name.
func ResolveNamed[T any](name string) (value T, found bool) {
	return (&GenericContainer[T]{Container: global}).ResolveNamed(name)
}

// ResolveAll will lookup and return all values registered.
func ResolveAll[T any]() (values []T) {
	return (&GenericContainer[T]{Container: global}).ResolveAll()
}
