package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return value should not be nil")
	} else {
		tracer.Trace("Testing trace...")
		if buf.String() != "Testing trace...\n" {
			t.Errorf("Trace should not write '%s'", buf.String())
		}
	}
}
