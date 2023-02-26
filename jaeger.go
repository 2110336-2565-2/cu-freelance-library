package gosdk

import (
	"go.opentelemetry.io/otel/sdk/trace"
	tr "go.opentelemetry.io/otel/trace"
)

type Jaeger interface {
	Tracer(name string, config ...tr.TracerOption) tr.Tracer
}

func NewJaeger(tracerProvider *trace.TracerProvider) Jaeger {
	return &jaeger{
		tracerProvider: tracerProvider,
	}
}

type jaeger struct {
	tracerProvider *trace.TracerProvider
}

func (j *jaeger) Tracer(name string, config ...tr.TracerOption) tr.Tracer {
	return j.tracerProvider.Tracer(name, config...)
}
