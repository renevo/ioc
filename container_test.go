package ioc_test

import (
	"context"
	"io"
	"testing"

	"github.com/renevo/ioc"
)

func TestContainer(t *testing.T) {
	ctx := ioc.WithContext(context.Background(), &ioc.Container{})

	ioc.RegisterToContext[io.Writer](ctx, &w{})
	ioc.RegisterNamedToContext[io.Writer](ctx, "discard", io.Discard)

	type writerFactory func() io.Writer

	ioc.RegisterToContext[writerFactory](ctx, func() io.Writer { return io.Discard })

	resolved, found := ioc.ResolveFromContext[io.Writer](ctx)
	if !found {
		t.Fatalf("Failed to resolve io.Writer")
	}

	t.Logf("Resolved io.Writer to %T", resolved)

	_, _ = resolved.Write([]byte("a"))

	resolved, found = ioc.ResolveNamedFromContext[io.Writer](ctx, "discard")
	if !found {
		t.Fatalf("Failed to resolve named io.Writer")
	}
	t.Logf("Resolved io.Writer to %T", resolved)
	_, _ = resolved.Write([]byte("a"))

	if _, found = ioc.ResolveFromContext[int](ctx); found {
		t.Logf("How did it find the int?")
	}

	resolvedFactory, found := ioc.ResolveFromContext[writerFactory](ctx)
	if !found {
		t.Fatalf("Failed to resolve named factory")
	}
	t.Logf("Resolved factory io.Writer to %T", resolvedFactory)
	_, _ = resolvedFactory().Write([]byte("b"))

	writers := ioc.ResolveAllFromContext[io.Writer](ctx)
	if len(writers) != 2 {
		t.Errorf("Failed to get all writers: expected 2; got %d", len(writers))
	}

	for _, writer := range writers {
		t.Logf("Found Writer: %T", writer)
	}
}

type w struct {
}

func (w) Write(p []byte) (n int, err error) { return len(p), nil }
