package ioc_test

import (
	"io"
	"testing"

	"github.com/renevo/ioc"
)

func TestContainer(t *testing.T) {
	ioc.Register[io.Writer](&w{})
	ioc.RegisterNamed("discard", io.Discard)

	resolved, found := ioc.Resolve[io.Writer]()
	if !found {
		t.Fatalf("Failed to resolve io.Writer")
	}

	t.Logf("Resolved io.Writer to %T", resolved)

	_, _ = resolved.Write([]byte("a"))

	resolved, found = ioc.ResolveNamed[io.Writer]("discard")
	if !found {
		t.Fatalf("Failed to resolve named io.Writer")
	}
	t.Logf("Resolved io.Writer to %T", resolved)

	_, _ = resolved.Write([]byte("a"))

	if _, found = ioc.Resolve[int](); found {
		t.Logf("How did it find the int?")
	}

	writers := ioc.ResolveAll[io.Writer]()
	if len(writers) != 2 {
		t.Errorf("Failed to get all writers: expected 2; got %d", len(writers))
	}

	for _, writer := range writers {
		t.Logf("Found Writer: %T", writer)
	}
}

type w struct {
}

func (w) Write(p []byte) (n int, err error) { return 0, nil }
