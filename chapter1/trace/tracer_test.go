package trace

import (
	"bytes"
	"golang.org/x/net/trace"
	"testing"
)

func TestNew(t *testing.T){
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		trace.Trace("Hello trace package.")
		if buf.String() != "Hello trace package.\n" {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}
	}
	t.Error("There's no test yet.")
}