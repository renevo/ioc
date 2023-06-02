package ioc_test

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/renevo/ioc"
)

func ExampleRegister() {
	ioc.Register(5)
	ioc.Register[io.Writer](os.Stdout)

	registered, _ := ioc.Resolve[int]()
	writer, _ := ioc.Resolve[io.Writer]()

	fmt.Fprintln(writer, registered)
	// output:
	// 5
}

func ExampleRegisterToContext() {
	ctx := ioc.WithContext(context.Background(), &ioc.Container{})

	ioc.RegisterToContext(ctx, 5)
	ioc.RegisterToContext[io.Writer](ctx, os.Stdout)

	registered, _ := ioc.ResolveFromContext[int](ctx)
	writer, _ := ioc.ResolveFromContext[io.Writer](ctx)

	fmt.Fprintln(writer, registered)
	// output:
	// 5
}
