package trace

import (
	"fmt"
	"io"
)

//Tracer used to define a tracing interface that takes zero or more arguments
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

type nilTracer struct{}

//New returns address of new tracer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

//Off API to stop/prevent tracing
func Off() Tracer {
	return &nilTracer{}
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

func (t *nilTracer) Trace(a ...interface{}) {}
