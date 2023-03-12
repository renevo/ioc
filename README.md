# Inversion Of Control Container

This is just a proof of concept IOC container in Go using generics.

With the way that generics work, the only way to really implement something like this is to have global functions, while the actual container is based on `reflect.Type` and `reflect.Value`.

Overall, the amount of code is pretty minimal since it doesn't do dependency injection. In theory the `Resolve` functions could be modified to not return founds and could be used to set things directly in *constructors*.
